package router

import (
	"database/sql"
	"hello/model"
	"time"

	"hello/firebase_setting" // ここは `firebase.go` のパスに置き換えてください

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func RouteAPI() {

	firebase_setting.Init()

	// DBの接続
	db = model.ConnectDB("sqlite3", "./model/Recode_verified.db")

	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // フロントエンドのURLを指定
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

	r.DELETE("/toilet/self", deleteSelfRecordHandler)

	r.PUT("/toilet/self", updateSelfRecordHandler)

	// r.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
	r.Run(":8080")
	defer db.Close()
}
