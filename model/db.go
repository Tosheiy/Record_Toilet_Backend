package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDB(driverName string, dataSourceName string) *sql.DB {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// InitDB()

	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func QueryDB(db *sql.DB, cmd string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(cmd, args...)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func InitDB() {
	// Read the SQL statements from the init.sql file
	sqlBytes, err := ioutil.ReadFile("./model/init.sql")
	if err != nil {
		log.Fatalf("Failed to read init.sql file: %v", err)
	}

	sqlStatements := string(sqlBytes)

	// テーブル作成実行
	fmt.Println(sqlStatements)
	_, err = db.Exec(sqlStatements)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
}
