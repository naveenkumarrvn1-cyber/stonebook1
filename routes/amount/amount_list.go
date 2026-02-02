package amount

import (
	"encoding/json"
	"net/http"
	"stonebook/constants"
)

type PaginationPayload struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type AmountResponse struct {
	AmountID int     `json:"amount_id"`
	Amount   float64 `json:"amount"`
}

func AmountList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var payload PaginationPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "Invalid JSON",
		})
		return
	}

	if payload.Limit == 0 {
		payload.Limit = 10
	}
	if payload.Page == 0 {
		payload.Page = 1
	}

	offset := (payload.Page - 1) * payload.Limit

	rows, err := constants.DB.Query("SELECT amount_id, amount FROM amount LIMIT ? OFFSET ?", payload.Limit, offset)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	defer rows.Close()

	var data []AmountResponse
	for rows.Next() {
		var a AmountResponse
		if err := rows.Scan(&a.AmountID, &a.Amount); err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": false,
				"error":  err.Error(),
			})
			return
		}
		data = append(data, a)
	}

	if data == nil {
		data = []AmountResponse{}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"data":   data,
		"page":   payload.Page,
		"limit":  payload.Limit,
	})
}
