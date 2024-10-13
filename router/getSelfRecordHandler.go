package router

import (
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSelfRecordHandler(c *gin.Context) {
	// Firebase認証
	VerifiedToken, err := verifiedAuth(c)
	if err != nil {
		log.Println("Error in verifiedAuth:", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// 特定のuidをuser_recordから取得
	new_user, err := model.CheckUserRecord(VerifiedToken.UID)
	if err == nil &&  new_user != nil{
		// 正常にユーザー情報を返す
		c.JSON(http.StatusOK, new_user)
	} else {
		// ユーザーが見つからなかった場合の処理
		log.Printf("error in CheckUserRecord or Non existed: %v", err)
		c.JSON(http.StatusOK, gin.H{"utid": "none", "apikey": "none"})
	}
}
