package router

import (
	"hello/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func postRecordHandler(c *gin.Context) {
	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var new_t_record model.TOILET_RECORD
	err = c.ShouldBindJSON(&new_t_record)
	if err != nil {
		log.Println("Error in ShouldBindJSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// 仕様を満たしているか確認
	err = checkRegulation(new_t_record)
	if err != nil {
		log.Println("Error in checkRegulation:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	new_t_record.ID = model.GenerateID() // uuidを生成
	new_t_record.Uid = VerifiedToken.UID

	// new_t_recordをtoilet_recordsに作成
	err = model.InsertRecordsDB(new_t_record)
	if err != nil {
		log.Println("Error in InsertRecordsDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	c.JSON(http.StatusCreated, new_t_record)
}
