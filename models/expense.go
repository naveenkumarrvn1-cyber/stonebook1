package models

type Expense struct {
	ExpenseId   int    `json:"expense_id"`
	ExpenseDate string `json:"expense_date"`
	Status      int    `json:"status"`
}
