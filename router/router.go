package router

import (
	"hello/firebase_setting"
	"hello/model"
	"log"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouteAPI(r *gin.Engine) {

	firebase_setting.Init()

	_, err := model.ConnectDB()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3306", "http://toiletrecordfront.s3-website-ap-northeast-1.amazonaws.com", "https://d2nehq3rb6dh1o.cloudfront.net"}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Record をすべて取得
	r.GET("/toilet", getAllRecordsHandler)

	// Record を挿入
	r.POST("/toilet", postRecordHandler)

	// id の Record を一つ取得
	r.GET("/toilet/:id", getOneRecordHandler)

	// id の Record の更新
	r.PUT("/toilet/:id", updateRecordHandler)

	// id の Record の削除
	r.DELETE("/toilet/:id", deleteRecordHandler)

	// id の Record の削除
	r.POST("/toilet/self/:utid", postSelfRecordHandler)

	// ユーザーのAuthorizationコード生成
	r.GET("/toilet/self/register", generateSelfRecordHandler)

	// ユーザーが登録したAuthorizationコードとUTIDを取得
	r.GET("/toilet/self", getSelfRecordHandler)

	// ユーザーが登録したAuthorizationコードとUTIDを削除
	r.DELETE("/toilet/self/:utid", deleteSelfRecordHandler)

	// ユーザーが登録したAuthorizationコードとUTIDを更新（削除して登録）
	r.PUT("/toilet/self/:utid", updateSelfRecordHandler)

	// デプロイテスト用
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Connection OK!!",
		})
	})
}
