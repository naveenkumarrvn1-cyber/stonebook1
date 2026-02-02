package product

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func ProductList(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 ProductList HANDLER STARTED")

	rows, err := constants.DB.Query(`
		SELECT
			product_id,
			product_name,
			status
		FROM product
		WHERE status = 1
		ORDER BY product_id DESC
	`)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {

		var (
			id     int
			name   string
			status int
		)

		if err := rows.Scan(&id, &name, &status); err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}

		products = append(products, map[string]interface{}{
			"product_id":   id,
			"product_name": name,
			"status":       status,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"count":  len(products),
		"data":   products,
	})
}
