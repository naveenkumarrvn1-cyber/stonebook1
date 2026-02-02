package bill

import (
	"encoding/json"
	"net/http"
	"time"

	"stonebook/constants"
)

/* ---------- Payload ---------- */
type PaginationPayload struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

/* ---------- Bill Response ---------- */
type BillResponse struct {
	LedgerType string `json:"ledger_type"`

	CompanyName string `json:"company_name"`
	ContactName string `json:"contact_name"`
	Phone       int64  `json:"phone"`

	ProductName string `json:"product_name"`

	ExpenseDate time.Time `json:"expense_date"`

	Amount float64 `json:"amount"`

	ReceiptDate time.Time `json:"receipt_date"`
}

/* ---------- BILL LIST API ---------- */
func BillList(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	/* ---------- Read payload (optional) ---------- */
	var payload PaginationPayload
	_ = json.NewDecoder(r.Body).Decode(&payload)

	/* ---------- Response object ---------- */
	var bill BillResponse

	/* ---------- 1. Ledger ---------- */
	constants.DB.QueryRow(`
		SELECT ledger_type
		FROM ledger
		WHERE status = 1
		ORDER BY id DESC
		LIMIT 1
	`).Scan(&bill.LedgerType)

	/* ---------- 2. Contact ---------- */
	constants.DB.QueryRow(`
		SELECT company_name, contact_name, phone
		FROM contact
		WHERE status = 1
		ORDER BY contact_id DESC
		LIMIT 1
	`).Scan(&bill.CompanyName, &bill.ContactName, &bill.Phone)

	/* ---------- 3. Product ---------- */
	constants.DB.QueryRow(`
		SELECT product_name
		FROM product
		WHERE status = 1
		ORDER BY product_id DESC
		LIMIT 1
	`).Scan(&bill.ProductName)

	/* ---------- 4. Expense ---------- */
	constants.DB.QueryRow(`
		SELECT expense_date
		FROM expense
		WHERE status = 1
		ORDER BY expense_id DESC
		LIMIT 1
	`).Scan(&bill.ExpenseDate)

	/* ---------- 5. Amount ---------- */
	constants.DB.QueryRow(`
		SELECT amount
		FROM amount
		ORDER BY amount_id DESC
		LIMIT 1
	`).Scan(&bill.Amount)

	/* ---------- 6. Receipt ---------- */
	constants.DB.QueryRow(`
		SELECT receipt_date
		FROM receipt
		WHERE status = 1
		ORDER BY receipt_id DESC
		LIMIT 1
	`).Scan(&bill.ReceiptDate)

	/* ---------- Final Response ---------- */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"data":   bill,
	})
}
