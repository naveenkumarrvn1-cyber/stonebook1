package conpay

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
	"stonebook/models"
)

func DeleteConPay(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 DeleteConPay HANDLER STARTED")

	var p models.ConPayDelete
	json.NewDecoder(r.Body).Decode(&p)

	if p.ContactID == 0 && p.PaymentID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "contact_id or payment_id is required",
		})
		return
	}

	if p.ContactID != 0 {
		constants.DB.Exec(
			`DELETE FROM contact WHERE contact_id = ?`,
			p.ContactID,
		)
	}

	if p.PaymentID != 0 {
		constants.DB.Exec(
			`DELETE FROM payment WHERE payment_id = ?`,
			p.PaymentID,
		)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Contact / Payment deleted successfully",
	})
}
