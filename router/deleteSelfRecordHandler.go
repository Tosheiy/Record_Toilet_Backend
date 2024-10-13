package router

import (
	"hello/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 現在の実装では存在しない場合もSuccessになる
func deleteSelfRecordHandler(c *gin.Context) {
	// paramの取得
	utid := c.Param("utid")

	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// 一致するuid と utid のUserRecordを削除
	err = model.DeleteUserRecord(VerifiedToken.UID, utid)
	if err != nil {
		log.Println("Error in DeleteUserRecord:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}