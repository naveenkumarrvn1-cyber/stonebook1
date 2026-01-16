package constants

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	// IMPORTANT: username = root, password = naveen
	DB, err = sql.Open(
		"mysql",
		"root:naveen@tcp(localhost:3306)/ledgerdb",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")
}
