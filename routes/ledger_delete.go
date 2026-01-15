package routes

import (
	"net/http"
	"strconv"

	"ledger/constants"
)

func LedgerDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ledgerID, _ := strconv.Atoi(id)

	_, err := constants.DB.Exec("DELETE FROM ledger WHERE ledger_id=?", ledgerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ledger Deleted Successfully"))
}
