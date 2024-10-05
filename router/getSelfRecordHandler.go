package router

import (
	"fmt"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getSelfRecordHandler(c *gin.Context) {

	// Authorizationヘッダーの取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Bearer プレフィックスを取り除く
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	idToken := parts[1]

	// トークンの検証
	VerifiedToken, err := firebase_setting.VerifyIDToken(idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Verified user id: %+v\n", VerifiedToken.UID)

	rows := model.QueryDB("SELECT * FROM user_table WHERE uid = ?", VerifiedToken.UID)
	defer rows.Close()

	var new_user model.USER

	if rows.Next() {
		err := rows.Scan(&new_user.ID, &new_user.UTID, &new_user.UID, &new_user.APIKEY)
		if err != nil {
			log.Fatal(err)
		}
		// 正常にユーザー情報を返す
		c.JSON(http.StatusOK, new_user)
	} else {
		// ユーザーが見つからなかった場合の処理
		c.JSON(http.StatusOK, gin.H{"utid": "none", "apikey": "none"})
	}
}
