package router

import (
	"fmt"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	// "time"
	"unicode/utf8"

	// "time"

	"github.com/gin-gonic/gin"
)

func updateRecordHandler(c *gin.Context) {
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
		fmt.Println("Error loading location:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Verified user id: %+v\n", VerifiedToken.UID)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error loading location:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
	}
	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		fmt.Println("Error loading location:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Created_at が既に設定されているかチェック
	if new_t_record.Created_at != "" {
		// 入力された日付のフォーマットが正しいかチェック
		_, err := time.Parse("2006-01-02 15:04", new_t_record.Created_at)
		if err != nil {
			fmt.Println("Error loading location:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Created_at が入力されていない場合は現在の時間を設定
		new_t_record.Created_at, err = CreateNowTime()
		if err != nil {
			fmt.Println("Error loading location:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Create Time Error."})
			return
		}
	}

	if utf8.RuneCountInString(new_t_record.Location) > 20 || utf8.RuneCountInString(new_t_record.Description) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文字数が多すぎます"})
		return
	}

	result, err := model.ExecDB("UPDATE toilet_records SET description = ?, length = ?, location = ?, feeling = ?, created_at = ? WHERE id = ? AND uid = ?",
		new_t_record.Description, new_t_record.Length, new_t_record.Location, new_t_record.Feeling, new_t_record.Created_at, id, VerifiedToken.UID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
}
