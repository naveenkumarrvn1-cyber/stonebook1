package payment

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 CreatePayment HANDLER STARTED")

	var p struct {
		PaymentDate string `json:"payment_date"`
	}

	json.NewDecoder(r.Body).Decode(&p)

	if p.PaymentDate == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "payment_date is required",
		})
		return
	}

	result, err := constants.DB.Exec(
		`INSERT INTO payment (payment_date, status) VALUES (?, 1)`,
		p.PaymentDate,
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
		"message":    "Payment created successfully",
		"payment_id": id,
	})
}
