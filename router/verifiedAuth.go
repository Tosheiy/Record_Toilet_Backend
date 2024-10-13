package router

import (
	"hello/firebase_setting"
	"net/http"
	"strings"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// verifiedAuth は Authorization ヘッダーを検証し、Firebase の ID トークンを返します
func verifiedAuth(c *gin.Context) (*auth.Token, error) {
	// Authorization ヘッダーの取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return nil, nil
	}

	// Bearer プレフィックスを取り除く
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return nil, nil
	}
	idToken := parts[1]

	// トークンの検証
	VerifiedToken, err := firebase_setting.VerifyIDToken(idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	return VerifiedToken, nil
}