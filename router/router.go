package router

import (
	"database/sql"
	"hello/model"
	"log"
	"os"
	"time"

	"hello/firebase_setting" // ここは `firebase.go` のパスに置き換えてください

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var db *sql.DB
var ginLambda *ginadapter.GinLambda

func RouteAPI() {

	firebase_setting.Init()

	// DBの接続
	// db = model.ConnectDB("sqlite3", "./model/Recode_verified.db")

	dsn := "admin:1049to@tcp(db:3306)/RecordToilet"
	db, err := model.ConnectDB("mysql", dsn)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}

	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3306"}, // フロントエンドのURLを指定
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


	// Lambda向けのアダプタを初期化
	ginLambda = ginadapter.New(r)

	// 環境を確認して処理を分岐
	if os.Getenv("IS_LAMBDA") == "" {
		// ローカル環境ではポート8080で起動
		r.Run(":8080")
	} else {
		// Lambda環境では、Lambda関数としてGinを起動
		lambda.Start(ginLambda.Proxy)
	}

	// r.RunTLS(":8080", "server.crt", "server.key")
	// r.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
	defer db.Close()
}
