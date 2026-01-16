package ledger

import (
	"encoding/json"
	"net/http"

	"ledger/constants"
	"ledger/models"
)

func LedgerCreate(w http.ResponseWriter, r *http.Request) {
	var ledger models.Ledger
	json.NewDecoder(r.Body).Decode(&ledger)

	query := `INSERT INTO ledger (ledger_name, ledger_type, ledger_description)
	          VALUES (?, ?, ?)`

	_, err := constants.DB.Exec(query, ledger.Name, ledger.Type, ledger.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ledger Created Successfully"))
}
