package router

import (
	"crypto/rand"
	"fmt"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func generateSelfRecordHandler(c *gin.Context) {

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

	UTID, err := generateKey(32)
	if err != nil {
		log.Fatalln(err)
	}
	APIKEY, err := generateKey(43)
	if err != nil {
		log.Fatalln(err)
	}
	APIKEY = "secret_" + APIKEY

	_, err = db.Exec("INSERT INTO user_table (uid, utid, apikey) VALUES (?, ?, ?)",
		VerifiedToken.UID, UTID, APIKEY)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User is exsited"})
		return
	}

	var new_user model.USER
	new_user.UTID = UTID
	new_user.APIKEY = APIKEY

	c.JSON(http.StatusCreated, new_user)
}

func generateKey(length int) (string, error) {
	// 使用する文字のセット（大文字、小文字、数字）
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, length)

	for i := range bytes {
		// 0からlen(charset)の範囲でランダムなインデックスを生成
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		bytes[i] = charset[index.Int64()]
	}

	return string(bytes), nil
}
