package contact

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"stonebook/constants"
)

/* ---------- Pagination Payload ---------- */
type PaginationPayload struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

/* ---------- Contact Response Struct (ORDER FIXED) ---------- */
type ContactResponse struct {
	ContactID   int    `json:"contact_id"`
	ContactName string `json:"contact_name"`

	ContactGroup int     `json:"contact_group"`
	Phone        int64   `json:"phone"`
	Email        *string `json:"email"`

	AddressLine1 *string `json:"address_line_1"`
	AddressLine2 *string `json:"address_line_2"`
	City         *string `json:"city"`
	State        *int    `json:"state"`
	Pincode      *int    `json:"pincode"`

	CompanyName *string `json:"company_name"`
	CompanyType string  `json:"company_type"`
	GstNo       *string `json:"gst_no"`

	OpeningBalanceType   string  `json:"opening_balance_type"`
	OpeningBalanceAmount float64 `json:"opening_balance_amount"`

	Tcs int `json:"tcs"`
	Tds int `json:"tds"`
	Rcm int `json:"rcm"`

	DiscountPercentage float64 `json:"discount_percentage"`
	CreditLimit        float64 `json:"credit_limit"`
	DueDays            int     `json:"due_days"`

	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

/* ---------- CONTACT LIST API ---------- */
func ContactList(w http.ResponseWriter, r *http.Request) {

	// CORS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("🔥 ContactList HANDLER STARTED")

	/* ---------- Default pagination ---------- */
	limit := 10
	page := 1

	/* ---------- Read JSON body (GET + payload) ---------- */
	if r.Body != nil {
		defer r.Body.Close()
		var payload PaginationPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err == nil {
			if payload.Limit > 0 {
				limit = payload.Limit
			}
			if payload.Page > 0 {
				page = payload.Page
			}
		}
	}

	offset := (page - 1) * limit
	log.Println("LIMIT:", limit, "PAGE:", page, "OFFSET:", offset)

	/* ---------- DB Query ---------- */
	rows, err := constants.DB.Query(`
		SELECT
			contact_id,
			contact_name,
			contact_group,
			phone,
			email,

			address_line_1,
			address_line_2,
			city,
			state,
			pincode,

			company_name,
			company_type,
			gst_no,

			opening_balance_type,
			opening_balance_amount,

			tcs,
			tds,
			rcm,

			discount_percentage,
			credit_limit,
			due_days,

			status,
			created_at
		FROM contact
		WHERE status = 1
		ORDER BY contact_id DESC
		LIMIT ? OFFSET ?
	`, limit, offset)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	defer rows.Close()

	var contacts []ContactResponse

	/* ---------- Scan rows ---------- */
	for rows.Next() {

		var c ContactResponse

		err := rows.Scan(
			&c.ContactID,
			&c.ContactName,
			&c.ContactGroup,
			&c.Phone,
			&c.Email,

			&c.AddressLine1,
			&c.AddressLine2,
			&c.City,
			&c.State,
			&c.Pincode,

			&c.CompanyName,
			&c.CompanyType,
			&c.GstNo,

			&c.OpeningBalanceType,
			&c.OpeningBalanceAmount,

			&c.Tcs,
			&c.Tds,
			&c.Rcm,

			&c.DiscountPercentage,
			&c.CreditLimit,
			&c.DueDays,

			&c.Status,
			&c.CreatedAt,
		)

		if err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}

		contacts = append(contacts, c)
	}

	/* ---------- Final Response ---------- */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"limit":  limit,
		"page":   page,
		"count":  len(contacts),
		"data":   contacts,
	})
}
