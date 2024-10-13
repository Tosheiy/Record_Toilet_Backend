package router

import (
	"database/sql"
	"hello/model"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func postSelfRecordHandler(c *gin.Context) {
	// paramの取得
	utid := c.Param("utid")

	/* postSelf認証 */
	userList, err := model.GetUserRecord(utid)
	if err != nil {
		if err == sql.ErrNoRows {
			// レコードが見つからない場合
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// その他のエラーの場合
		log.Println("Error in GetUserRecord:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	user := userList[0]

	// Authorizationヘッダーの取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Bearer プレフィックスを取り除く
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	idToken := parts[1]

	// API キーの検証
	if idToken != user.APIKEY {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		c.Abort()
		return
	}

	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		log.Println("Error in ShouldBindJSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// 仕様を満たしているか確認
	err = checkRegulation(new_t_record)
	if err != nil {
		log.Println("Error in checkRegulation:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	new_t_record.ID = model.GenerateID() // uuidの生成
	new_t_record.Uid = user.APIKEY

	// new_t_record をtoilet_recordsに作成
	err = model.InsertRecordsDB(new_t_record)
	if err != nil {
		log.Println("Error in InsertRecordsDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, new_t_record)
}
