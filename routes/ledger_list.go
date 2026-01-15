package routes

import (
	"encoding/json"
	"net/http"

	"ledger/constants"
	"ledger/models"
)

func LedgerList(w http.ResponseWriter, r *http.Request) {
	rows, err := constants.DB.Query("SELECT ledger_id, ledger_name, ledger_type, ledger_description FROM ledger")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ledgers []models.Ledger

	for rows.Next() {
		var l models.Ledger
		rows.Scan(&l.LedgerID, &l.LedgerName, &l.LedgerType, &l.LedgerDescription)
		ledgers = append(ledgers, l)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ledgers)
}
