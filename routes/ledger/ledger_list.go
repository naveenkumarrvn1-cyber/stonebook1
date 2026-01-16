package ledger

import (
	"encoding/json"
	"net/http"

	"ledger/constants"
	"ledger/models"
)

func GetLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := constants.DB.Query(
		"SELECT id, name, type, description FROM ledger",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	ledgers := []models.Ledger{}

	for rows.Next() {
		var l models.Ledger
		if err := rows.Scan(&l.ID, &l.Name, &l.Type, &l.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ledgers = append(ledgers, l)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ledgers)
}
