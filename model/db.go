package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDB(driverName string, dataSourceName string) (*sql.DB, error) {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 接続が実際に確立されているか確認
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// InitDB()

	return db, nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func QueryDB(cmd string, args ...interface{}) *sql.Rows {

	rows, err := db.Query(cmd, args...)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}


func ExecDB(cmd string, args ...interface{}) (sql.Result, error) {

	result, err := db.Exec(cmd, args...)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
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
