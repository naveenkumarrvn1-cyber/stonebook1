package models

type Contact struct {
	ContactID int `json:"contact_id"`

	ContactName  string `json:"contact_name"`
	ContactGroup int    `json:"contact_group"`

	Phone int64  `json:"phone"`
	Email string `json:"email"`

	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        int    `json:"state"`
	Pincode      int    `json:"pincode"`

	CompanyName string `json:"company_name"`
	CompanyType string `json:"company_type"`

	GSTNo string `json:"gst_no"`

	OpeningBalanceType   string  `json:"opening_balance_type"`
	OpeningBalanceAmount float64 `json:"opening_balance_amount"`

	TCS int `json:"tcs"`
	TDS int `json:"tds"`
	RCM int `json:"rcm"`

	DiscountPercentage float64 `json:"discount_percentage"`
	CreditLimit        float64 `json:"credit_limit"`
	DueDays            int     `json:"due_days"`

	Status int `json:"status"`
}
