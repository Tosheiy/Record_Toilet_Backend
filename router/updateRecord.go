package router

import (
	"hello/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func updateRecordHandler(c *gin.Context) {
	// パラメーターの取得
	id := c.Param("id")

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

	new_t_record.ID = id
	new_t_record.Uid = VerifiedToken.UID

	// 特定の要素をupdateする
	err = model.UpdateRecordDB(new_t_record.Uid, new_t_record.ID, new_t_record)
	if err != nil {
		log.Println("Error in UpdateRecordDB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"success": "Create"})
}
