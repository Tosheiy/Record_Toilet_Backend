package router

import (
	"database/sql"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"
	"strings"
	"time"

	// ここは `firebase.go` のパスに置き換えてください

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func RouteAPI(r *gin.Engine) {

	firebase_setting.Init()

	// DBの接続
	// db = model.ConnectDB("sqlite3", "./model/Recode_verified.db")

    // dbEndpoint := "mytoiletrecord.cpwiqm8ec2pr.ap-northeast-1.rds.amazonaws.com:3306"
    // dbUser := "admin"
    // dbPassword := "1049to1049ad"
    // dbName := "mytoiletrecord"

    // dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, dbName)

    // dbHost := os.Getenv("DB_HOST")
    // dbUser := os.Getenv("DB_USER")
    // dbPassword := os.Getenv("DB_PASSWORD")
    // dbName := os.Getenv("DB_NAME")

	// log.Print(dbUser) //中身空白

    // // RDS Proxy を通じてデータベースに接続
    // dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	dsn := "admin:1049to1049ad@tcp(mytoiletrecord.cpwiqm8ec2pr.ap-northeast-1.rds.amazonaws.com:3306)/mytoiletrecord"
	
	db, err := model.ConnectDB("mysql", dsn)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3306", "http://toiletrecordfront.s3-website-ap-northeast-1.amazonaws.com"}, // フロントエンドのURLを指定
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

	// デプロイテスト用
	r.GET("/ping", func(c *gin.Context) {
		// Authorizationヘッダーの取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// Bearer プレフィックスを取り除く
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		idToken := parts[1]

		// トークンの検証
		VerifiedToken, err := firebase_setting.VerifyIDToken(idToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": VerifiedToken.UID,
		})
	})

	// // 環境を確認して処理を分岐
	// if os.Getenv("IS_LAMBDA") == "" {
	// 	// ローカル環境ではポート8080で起動
	// 	r.Run(":8080")
	// } else {
	// 	// Lambda環境では、Lambda関数としてGinを起動
	// 	lambda.Start(ginLambda.Proxy)
	// }

	// r.RunTLS(":8080", "server.crt", "server.key")
	// r.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
	// defer db.Close()
}
