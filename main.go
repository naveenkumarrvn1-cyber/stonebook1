package main

import (
	"log"
	"net/http"

	"ledger/constants"
	"ledger/routes"
)

func main() {
	constants.ConnectDB()

	http.HandleFunc("/ledger/list", routes.LedgerList)
	http.HandleFunc("/ledger/create", routes.LedgerCreate)
	http.HandleFunc("/ledger/update", routes.LedgerUpdate)
	http.HandleFunc("/ledger/delete", routes.LedgerDelete)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
