package ledger

import (
	"encoding/json"
	"net/http"

	"ledger/constants"
	"ledger/models"
)

func CreateLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var l models.Ledger
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := constants.DB.Exec(
		"INSERT INTO ledger (name, type, description) VALUES (?, ?, ?)",
		l.Name, l.Type, l.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	l.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(l)
}
