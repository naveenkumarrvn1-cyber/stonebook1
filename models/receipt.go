package models

type Receipt struct {
	ReceiptId   int    `json:"receipt_id"`
	ReceiptDate string `json:"receipt_date"`
	Status      int    `json:"status"`
}
