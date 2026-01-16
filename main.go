package main

import (
	"log"
	"net/http"

	"ledger/constants"
	"ledger/middleware"
	contact "ledger/routes/contact"
	ledger "ledger/routes/ledger"

	"github.com/gorilla/mux"
)

func main() {

	// Connect DB
	constants.ConnectDB()

	// Create router
	r := mux.NewRouter()

	// =====================
	// Ledger Routes
	// =====================
	r.HandleFunc("/ledger", ledger.GetLedger).Methods("GET")
	r.HandleFunc("/ledger", ledger.CreateLedger).Methods("POST")
	r.HandleFunc("/ledger/{id}", ledger.UpdateLedger).Methods("PUT")
	r.HandleFunc("/ledger/{id}", ledger.DeleteLedger).Methods("DELETE")

	// =====================
	// Contact Routes
	// =====================
	r.HandleFunc("/contacts", contact.ContactList).Methods("GET")
	r.HandleFunc("/contacts/{id}", contact.ContactEdit).Methods("GET")
	r.HandleFunc("/contacts/{id}", contact.ContactUpdate).Methods("PUT")
	r.HandleFunc("/contacts/{id}", contact.ContactDelete).Methods("DELETE")

	// Enable CORS
	handler := middleware.EnableCORS(r)

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
