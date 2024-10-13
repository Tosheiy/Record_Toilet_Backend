package router

import (
	"hello/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func updateSelfRecordHandler(c *gin.Context) {
	//　パラメーターの取得
	utid := c.Param("utid")

	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// 削除して追加するという方式
	// 削除
	err = model.DeleteUserRecord(VerifiedToken.UID, utid)
	if err != nil {
		log.Println("Error in DeleteUserRecord:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	// 作成
	UTID, err := generateKey(32)
	if err != nil {
		log.Println("Error in generateKey:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	APIKEY, err := generateKey(43)
	if err != nil {
		log.Println("Error in generateKey:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	APIKEY = "secret_" + APIKEY

	var new_user model.USER
	new_user.UID = VerifiedToken.UID
	new_user.UTID = UTID
	new_user.APIKEY = APIKEY

	// UserRecordに作成
	err = model.InsertUserRecord(new_user)
	if err != nil {
		log.Println("Error in InsertUserRecord:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User is exsited"})
		return
	}

	c.JSON(http.StatusCreated, new_user)
}
