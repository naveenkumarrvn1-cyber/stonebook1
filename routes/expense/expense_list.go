package expense

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"stonebook/constants"
)

/* ---------- Pagination Payload ---------- */
type PaginationPayload struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

/* ---------- Expense Response Struct (ORDER GUARANTEED) ---------- */
type ExpenseResponse struct {
	ExpenseID   int       `json:"expense_id"`
	ExpenseDate time.Time `json:"expense_date"`
	Status      int       `json:"status"`
}

/* ---------- EXPENSE LIST API ---------- */
func ExpenseList(w http.ResponseWriter, r *http.Request) {

	// CORS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 ExpenseList HANDLER STARTED")

	/* ---------- Default pagination ---------- */
	limit := 10
	page := 1

	/* ---------- Read JSON body (GET + payload) ---------- */
	if r.Body != nil {
		defer r.Body.Close()

		var payload PaginationPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err == nil {
			if payload.Limit > 0 {
				limit = payload.Limit
			}
			if payload.Page > 0 {
				page = payload.Page
			}
		}
	}

	offset := (page - 1) * limit
	log.Println("LIMIT:", limit, "PAGE:", page, "OFFSET:", offset)

	/* ---------- DB Query ---------- */
	rows, err := constants.DB.Query(`
		SELECT
			expense_id,
			expense_date,
			status
		FROM expense
		WHERE status = 1
		ORDER BY expense_id DESC
		LIMIT ? OFFSET ?
	`, limit, offset)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	defer rows.Close()

	var expenses []ExpenseResponse

	/* ---------- Scan rows ---------- */
	for rows.Next() {

		var e ExpenseResponse

		err := rows.Scan(
			&e.ExpenseID,
			&e.ExpenseDate,
			&e.Status,
		)
		if err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}

		expenses = append(expenses, e)
	}

	/* ---------- Final response ---------- */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"limit":  limit,
		"page":   page,
		"count":  len(expenses),
		"data":   expenses,
	})
}
