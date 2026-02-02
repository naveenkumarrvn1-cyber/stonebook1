package receipt

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func CreateReceipt(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 CreateReceipt HANDLER STARTED")

	var p struct {
		ReceiptDate string `json:"receipt_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "Invalid JSON",
		})
		return
	}

	if p.ReceiptDate == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "receipt_date is required",
		})
		return
	}

	result, err := constants.DB.Exec(
		`INSERT INTO receipt (receipt_date, status) VALUES (?, 1)`,
		p.ReceiptDate,
	)

	if err != nil {
		log.Println("❌ DB ERROR:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     true,
		"message":    "Receipt created successfully",
		"receipt_id": id,
	})
}
