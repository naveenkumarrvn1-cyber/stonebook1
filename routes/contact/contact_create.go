package contact

import (
	"encoding/json"
	"log"
	"net/http"

	"stonebook/constants"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 CreateContact HANDLER STARTED")

	var p struct {
		ContactName  string `json:"contact_name"`
		ContactGroup int    `json:"contact_group"`
		Phone        int64  `json:"phone"`
		Email        string `json:"email"`

		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		City         string `json:"city"`
		State        int    `json:"state"`
		Pincode      int    `json:"pincode"`

		CompanyName string `json:"company_name"`
		CompanyType string `json:"company_type"`
		GSTNo       string `json:"gst_no"`

		OpeningBalanceType   string  `json:"opening_balance_type"`
		OpeningBalanceAmount float64 `json:"opening_balance_amount"`

		TCS int `json:"tcs"`
		TDS int `json:"tds"`
		RCM int `json:"rcm"`

		DiscountPercentage float64 `json:"discount_percentage"`
		CreditLimit        float64 `json:"credit_limit"`
		DueDays            int     `json:"due_days"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "Invalid JSON",
		})
		return
	}

	// Required validations
	if p.ContactName == "" || p.Phone == 0 || p.ContactGroup == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  "contact_name, contact_group, phone are required",
		})
		return
	}

	result, err := constants.DB.Exec(`
		INSERT INTO contact (
			contact_name, contact_group, phone, email,
			address_line_1, address_line_2, city, state, pincode,
			company_name, company_type, gst_no,
			opening_balance_type, opening_balance_amount,
			tcs, tds, rcm,
			discount_percentage, credit_limit, due_days,
			status
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)
	`,
		p.ContactName,
		p.ContactGroup,
		p.Phone,
		nullIfEmpty(p.Email),

		nullIfEmpty(p.AddressLine1),
		nullIfEmpty(p.AddressLine2),
		nullIfEmpty(p.City),
		intOrNull(p.State),
		intOrNull(p.Pincode),

		nullIfEmpty(p.CompanyName),
		defaultCompanyType(p.CompanyType),
		nullIfEmpty(p.GSTNo),

		defaultOBType(p.OpeningBalanceType),
		p.OpeningBalanceAmount,

		p.TCS,
		p.TDS,
		p.RCM,

		p.DiscountPercentage,
		p.CreditLimit,
		p.DueDays,
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
		"message":    "Contact created successfully",
		"contact_id": id,
	})
}

/* helpers */

func nullIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

func intOrNull(v int) interface{} {
	if v == 0 {
		return nil
	}
	return v
}

func defaultCompanyType(v string) string {
	if v == "" {
		return "unregistered"
	}
	return v
}

func defaultOBType(v string) string {
	if v == "" {
		return "CR"
	}
	return v
}
