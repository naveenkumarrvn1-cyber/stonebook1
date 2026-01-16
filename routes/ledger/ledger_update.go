package ledger

import (
	"encoding/json"
	"net/http"

	"ledger/constants"
	"ledger/models"

	"github.com/gorilla/mux"
)

func UpdateLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var l models.Ledger
	json.NewDecoder(r.Body).Decode(&l)

	_, err := constants.DB.Exec(
		"UPDATE ledger SET name=?, type=?, description=? WHERE id=?",
		l.Name, l.Type, l.Description, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Ledger updated",
	})
}
