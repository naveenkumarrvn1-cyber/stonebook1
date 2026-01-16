package ledger

import (
	"encoding/json"
	"net/http"

	"ledger/constants"

	"github.com/gorilla/mux"
)

func DeleteLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	_, err := constants.DB.Exec(
		"DELETE FROM ledger WHERE id=?",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Ledger deleted",
	})
}
