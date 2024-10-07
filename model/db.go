package model

import (
	"database/sql"
	"fmt"
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

	InitDB()

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
	// sqlBytes, err := os.ReadFile("./init/init.sql")
	// if err != nil {
	// 	log.Fatalf("Failed to read init.sql file: %v", err)
	// }

	// sqlStatements := string(sqlBytes)

	// // テーブル作成実行
	// fmt.Println(sqlStatements)
	// _, err = db.Exec(sqlStatements)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully!")

    // Create database
    _, err := db.Exec("CREATE DATABASE IF NOT EXISTS RecordToilet;")
    if err != nil {
        log.Fatalf("Error creating database: %v", err)
    }

    // Use the database
    _, err = db.Exec("USE RecordToilet;")
    if err != nil {
        log.Fatalf("Error selecting database: %v", err)
    }

    // Create tables
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS toilet_records (
        id INT AUTO_INCREMENT PRIMARY KEY,
        description TEXT,
        created_at DATETIME,
        length INT,
        location TEXT,
        feeling INT,
        uid VARCHAR(255)
    );`)
    if err != nil {
        log.Fatalf("Error creating table toilet_records: %v", err)
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS user_table (
        id INT AUTO_INCREMENT PRIMARY KEY,
        utid VARCHAR(32) UNIQUE NOT NULL,
        uid VARCHAR(255) UNIQUE NOT NULL,
        apikey VARCHAR(50) UNIQUE NOT NULL
    );`)
    if err != nil {
        log.Fatalf("Error creating table user_table: %v", err)
    }

    fmt.Println("Tables created successfully!")
}
