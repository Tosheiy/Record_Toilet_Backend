package model

import (
	"database/sql"
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