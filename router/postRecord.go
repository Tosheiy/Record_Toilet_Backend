package router

import (
	"fmt"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func postRecordHandler(c *gin.Context) {
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
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Verified user id: %+v\n", VerifiedToken.UID)

	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Created_at が既に設定されているかチェック
	if new_t_record.Created_at != "" {
		// 入力された日付のフォーマットが正しいかチェック
		_, err := time.Parse("2006-01-02 15:04", new_t_record.Created_at)
		if err != nil {
			log.Printf("JSON binding error: %v", err) // エラーメッセージをログに出力
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Expected format: 2006-01-02 15:04"})
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

	// 文字数制限確認
	if utf8.RuneCountInString(new_t_record.Location) > 20 || utf8.RuneCountInString(new_t_record.Description) > 50 {
		log.Printf("JSON binding error: %v", err) // エラーメッセージをログに出力
		c.JSON(http.StatusBadRequest, gin.H{"error": "文字数が多すぎます"})
		return
	}


	_, err = model.ExecDB("INSERT INTO toilet_records (description, created_at, length, location, feeling, uid) VALUES (?, ?, ?, ?, ?, ?)",
		new_t_record.Description, new_t_record.Created_at, new_t_record.Length, new_t_record.Location, new_t_record.Feeling, VerifiedToken.UID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DB挿入時にエラーが発生しました"})
		return
	}
	
	c.JSON(http.StatusCreated, new_t_record)
}
