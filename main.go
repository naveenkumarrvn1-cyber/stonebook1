package main

import (
	"log"
	"net/http"
	"stonebook/constants"
	"stonebook/middleware"

	"stonebook/routes/amount"
	"stonebook/routes/bill"
	"stonebook/routes/conpay"
	"stonebook/routes/contact"
	"stonebook/routes/expense"
	"stonebook/routes/ledger"
	"stonebook/routes/payment"
	"stonebook/routes/product"
	"stonebook/routes/receipt"
)

func main() {
	constants.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/receipt/create", receipt.CreateReceipt)
	mux.HandleFunc("/receipt/delete", receipt.DeleteReceipt)
	mux.HandleFunc("/receipt/list", receipt.ReceiptList)

	mux.HandleFunc("/payment/create", payment.CreatePayment)
	mux.HandleFunc("/payment/delete", payment.DeletePayment)
	mux.HandleFunc("/payment/list", payment.PaymentList)

	mux.HandleFunc("/expense/create", expense.CreateExpense)
	mux.HandleFunc("/expense/delete", expense.DeleteExpense)
	mux.HandleFunc("/expense/list", expense.ExpenseList)

	mux.HandleFunc("/product/create", product.CreateProduct)
	mux.HandleFunc("/product/delete", product.DeleteProduct)
	mux.HandleFunc("/product/list", product.ProductList)

	mux.HandleFunc("/contact/create", contact.CreateContact)
	mux.HandleFunc("/contact/delete", contact.DeleteContact)
	mux.HandleFunc("/contact/list", contact.ContactList)

	mux.HandleFunc("/ledger/create", ledger.CreateLedger)
	mux.HandleFunc("/ledger/delete", ledger.DeleteLedger)
	mux.HandleFunc("/ledger/list", ledger.LedgerList)

	mux.HandleFunc("/amount/list", amount.AmountList)

	mux.HandleFunc("/conpay/delete", conpay.DeleteConPay)

	mux.HandleFunc("/bill/list", bill.BillList)

	log.Fatal(http.ListenAndServe(":8080", middleware.EnableCORS(mux)))

}
