package router

import (
	"database/sql"
	"fmt"
	"hello/model"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func postSelfRecordHandler(c *gin.Context) {
	// paramの取得
	utid := c.Param("utid")

	var user model.USER
	err := db.QueryRow("SELECT * FROM user_table WHERE utid = ?", utid).Scan(
		&user.ID,
		&user.UTID,
		&user.UID,
		&user.APIKEY,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// レコードが見つからない場合
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// その他のエラーの場合
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed:"})
		return
	}

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

	// API キーの検証
	if idToken != user.APIKEY {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		c.Abort()
		return
	}

	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Created_at が既に設定されているかチェック
	if new_t_record.Created_at != "" {
		// 入力された日付のフォーマットが正しいかチェック
		_, err := time.Parse("2006-01-02 15:04", new_t_record.Created_at)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Expected format: 2006-01-02T15:04"})
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

	_, err = model.ExecDB("INSERT INTO toilet_records (description, created_at, length, location, feeling, uid) VALUES (?, ?, ?, ?, ?, ?)",
		new_t_record.Description, new_t_record.Created_at, new_t_record.Length, new_t_record.Location, new_t_record.Feeling, user.UID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DB挿入時にエラーが発生しました"})
		return
	}

	c.JSON(http.StatusCreated, new_t_record)
}