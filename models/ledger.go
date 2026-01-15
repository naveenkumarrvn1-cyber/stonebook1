package models

type Ledger struct {
	LedgerID          int    `json:"ledger_id"`
	LedgerName        string `json:"ledger_name"`
	LedgerType        string `json:"ledger_type"`
	LedgerDescription string `json:"ledger_description"`
}
