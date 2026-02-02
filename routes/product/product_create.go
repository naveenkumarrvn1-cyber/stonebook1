package product

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("🔥 CreateProduct HANDLER STARTED")

	// CORS preflight
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 CreateProduct HANDLER STARTED")

	// Payload
	var p struct {
		ProductName string `json:"product_name"`
	}

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println(" JSON DECODE ERROR:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "Invalid JSON payload",
		})
		return
	}

	if p.ProductName == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "product_name is required",
		})
		return
	}

	// Insert DB
	result, err := constants.DB.Exec(
		`INSERT INTO product (product_name, status) VALUES (?, ?)`,
		p.ProductName,
		1,
	)

	if err != nil {
		log.Println(" DB ERROR:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Product created successfully",
		"id":      id,
	})
}
