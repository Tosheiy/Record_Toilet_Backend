package router

import (
	"database/sql"
	"hello/model"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func RouteAPI() {
	db = model.ConnectDB("sqlite3", "./model/Recode.db")

	r := gin.Default()

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

	r.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
	defer db.Close()
}
