package receipt

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

/* ---------- Receipt Response Struct (ORDER GUARANTEED) ---------- */
type ReceiptResponse struct {
	ReceiptID   int       `json:"receipt_id"`
	ReceiptDate time.Time `json:"receipt_date"`
	Status      int       `json:"status"`
}

/* ---------- RECEIPT LIST API ---------- */
func ReceiptList(w http.ResponseWriter, r *http.Request) {

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
	log.Println("🔥 ReceiptList HANDLER STARTED")

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
			receipt_id,
			receipt_date,
			status
		FROM receipt
		WHERE status = 1
		ORDER BY receipt_id DESC
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

	var receipts []ReceiptResponse

	/* ---------- Scan rows ---------- */
	for rows.Next() {

		var rcp ReceiptResponse

		err := rows.Scan(
			&rcp.ReceiptID,
			&rcp.ReceiptDate,
			&rcp.Status,
		)
		if err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}

		receipts = append(receipts, rcp)
	}

	/* ---------- Final response ---------- */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"limit":  limit,
		"page":   page,
		"count":  len(receipts),
		"data":   receipts,
	})
}

