package router

import (
	"hello/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func deleteRecordHandler(c *gin.Context) {
	// パラメータの読み取り
	id := c.Param("id")

	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// uid と id が一致したRecordを削除
	err = model.DeleteRecordsDB(VerifiedToken.UID, id)
	if err != nil {
		log.Println("Error in DeleteRecordsDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Delete"})
}
