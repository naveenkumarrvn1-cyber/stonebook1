package constants

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/stonebook?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 🔒 FORCE SINGLE DB POOL
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	DB = db

	// 🔥 FINAL PROOF (keep this always)
	var dbName string
	db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	log.Println("✅ CONNECTED TO DB:", dbName)
}
