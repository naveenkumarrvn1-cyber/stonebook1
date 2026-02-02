package expense

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func CreateExpense(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 CreateExpense HANDLER STARTED")

	var p struct {
		ExpenseDate string `json:"expense_date"`
	}

	json.NewDecoder(r.Body).Decode(&p)

	if p.ExpenseDate == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "expense_date is required",
		})
		return
	}

	result, err := constants.DB.Exec(
		`INSERT INTO expense (expense_date, status) VALUES (?, 1)`,
		p.ExpenseDate,
	)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     true,
		"message":    "Expense created successfully",
		"expense_id": id,
	})
}
