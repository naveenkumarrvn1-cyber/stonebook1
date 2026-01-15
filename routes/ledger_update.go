package routes

import (
	"net/http"
	"strconv"

	"ledger/constants"
)

func LedgerUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	query := `UPDATE ledger SET ledger_name=?, ledger_type=?, ledger_description=? WHERE ledger_id=?`

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ledgerID, _ := strconv.Atoi(id)

	_, err = constants.DB.Exec(
		query,
		r.FormValue("ledger_name"),
		r.FormValue("ledger_type"),
		r.FormValue("ledger_description"),
		ledgerID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ledger Updated Successfully"))
}
