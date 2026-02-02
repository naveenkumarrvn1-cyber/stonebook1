package models

type Payment struct {
	PaymentId   int    `json:"payment_id"`
	PaymentDate string `json:"payment_date"`
	Status      int    `json:"status"`
}
