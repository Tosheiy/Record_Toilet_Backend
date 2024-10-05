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

func deleteSelfRecordHandler(c *gin.Context) {

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

	result, err := model.ExecDB("DELETE FROM user_table WHERE uid = ?", VerifiedToken.UID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	c.JSON(http.StatusOK, result)
}