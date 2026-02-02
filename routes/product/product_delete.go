package product

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 DeleteProduct HANDLER STARTED")

	var p struct {
		ProductID int `json:"product_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "Invalid JSON",
		})
		return
	}

	if p.ProductID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "product_id is required",
		})
		return
	}

	result, err := constants.DB.Exec(
		`DELETE FROM product WHERE product_id = ?`,
		p.ProductID,
	)

	if err != nil {
		log.Println("❌ DB ERROR:", err)
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
			"error":  "No product found",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Product deleted successfully",
	})
}
