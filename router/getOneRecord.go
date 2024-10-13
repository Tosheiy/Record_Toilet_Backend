package router

import (
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOneRecordHandler(c *gin.Context) {
	// パラメータの取得
	id := c.Param("id")

	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// 特定の uid と id のレコードを取得
	toilet_record, err := model.GetOneRecordDB(VerifiedToken.UID, id)
	if err != nil {
		log.Println("Error in GetOneRecordDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, toilet_record)
}