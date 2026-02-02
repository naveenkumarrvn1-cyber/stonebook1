package payment

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func DeletePayment(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 DeletePayment HANDLER STARTED")

	var p struct {
		PaymentID int `json:"payment_id"`
	}

	json.NewDecoder(r.Body).Decode(&p)

	if p.PaymentID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "payment_id is required",
		})
		return
	}

	constants.DB.Exec(
		`DELETE FROM payment WHERE payment_id = ?`,
		p.PaymentID,
	)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Payment deleted successfully",
	})
}
