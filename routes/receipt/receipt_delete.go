package receipt

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func DeleteReceipt(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 DeleteReceipt HANDLER STARTED")

	var p struct {
		ReceiptID int `json:"receipt_id"`
	}

	json.NewDecoder(r.Body).Decode(&p)

	if p.ReceiptID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "receipt_id is required",
		})
		return
	}

	result, err := constants.DB.Exec(
		`DELETE FROM receipt WHERE receipt_id = ?`,
		p.ReceiptID,
	)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "No receipt found",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Receipt deleted successfully",
	})
}
