package expense

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func DeleteExpense(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 DeleteExpense HANDLER STARTED")

	var p struct {
		ExpenseID int `json:"expense_id"`
	}

	json.NewDecoder(r.Body).Decode(&p)

	if p.ExpenseID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "expense_id is required",
		})
		return
	}

	constants.DB.Exec(
		`DELETE FROM expense WHERE expense_id = ?`,
		p.ExpenseID,
	)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Expense deleted successfully",
	})
}
