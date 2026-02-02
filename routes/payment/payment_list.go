package payment

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

/* ---------- Payment Response Struct (ORDER GUARANTEED) ---------- */
type PaymentResponse struct {
	PaymentID   int       `json:"payment_id"`
	PaymentDate time.Time `json:"payment_date"`
	Status      int       `json:"status"`
}

/* ---------- PAYMENT LIST API ---------- */
func PaymentList(w http.ResponseWriter, r *http.Request) {

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
	log.Println("🔥 PaymentList HANDLER STARTED")

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
			payment_id,
			payment_date,
			status
		FROM payment
		WHERE status = 1
		ORDER BY payment_id DESC
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

	var payments []PaymentResponse

	/* ---------- Scan rows ---------- */
	for rows.Next() {

		var p PaymentResponse

		err := rows.Scan(
			&p.PaymentID,
			&p.PaymentDate,
			&p.Status,
		)
		if err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}

		payments = append(payments, p)
	}

	/* ---------- Final response ---------- */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"limit":  limit,
		"page":   page,
		"count":  len(payments),
		"data":   payments,
	})
}
