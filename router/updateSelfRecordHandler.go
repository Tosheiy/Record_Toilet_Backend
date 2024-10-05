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

func updateSelfRecordHandler(c *gin.Context) {

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

	_, err = model.ExecDB("UPDATE user_table SET utid = ?, apikey = ? WHERE uid = ?",
		UTID, APIKEY, VerifiedToken.UID)
	if err != nil {
		log.Fatal(err)
	}

	var new_user model.USER
	new_user.UTID = UTID
	new_user.APIKEY = APIKEY

	c.JSON(http.StatusCreated, new_user)
}
