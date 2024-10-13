package router

import (
	// "fmt"
	// "hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "strings"
)

func getAllRecordsHandler(c *gin.Context) {
	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}
	
	// 特定のUIDのレコードを全て取得
	toilet_records, err := model.GetAllRecordsDB(VerifiedToken.UID)
	if err != nil {
		log.Println("Error in GetAllRecordsDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, toilet_records)
}
