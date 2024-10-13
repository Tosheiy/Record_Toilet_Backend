package router

import (
	"crypto/rand"
	"fmt"

	"hello/model"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateSelfRecordHandler(c *gin.Context) {
	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// 同じUIDがテーブルに存在しないことを確認
	existedUser, _ := model.CheckUserRecord(VerifiedToken.UID)
	if existedUser != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User is exsited"})
		return
	}

	// selfPost用UTIDを作成（32文字）
	UTID, err := generateKey(32)
	if err != nil {
		log.Println("Error in generateKey:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	// selfPost用APIKEYを作成（43文字＋7文字）
	APIKEY, err := generateKey(43)
	if err != nil {
		log.Println("Error in generateKey:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	APIKEY = "secret_" + APIKEY //secret_接頭語を作成


	var new_user model.USER
	new_user.UID = VerifiedToken.UID
	new_user.UTID = UTID
	new_user.APIKEY = APIKEY

	// userをuser_tableに追加
	err = model.InsertUserRecord(new_user)
	if err != nil {
		log.Println("Error in InsertUserRecord:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

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
