## StoneBook Backend API

A backend REST API application built using Golang and MySQL.

## Tech Stack
- Golang
- MySQL
- REST API

## Features
- CRUD operations for Contacts
- Product management
- Payment & Expense tracking
- JSON based APIs
- CORS enabled

## How to Run
1. Clone the repository
2. Update MySQL credentials in config file
3. Run `go run main.go`
4. Test APIs using Postman



    ## API Endpoints

### Contact
- POST /contact/create
- POST /contact/delete
- GET  /contact/list

### Product
- POST /product/create
- POST /product/delete
- GET  /product/list

### Payment
- POST /payment/create
- POST /payment/delete
- GET  /payment/list

### Expense
- POST /expense/create
- POST /expense/delete
- GET  /expense/list

### Receipt
- POST /receipt/create
- POST /receipt/delete
- GET  /receipt/list

### Ledger
- POST /ledger/create
- POST /ledger/delete
- GET  /ledger/list

### Others
- GET  /amount/list
- POST /conpay/delete
- GET  /bill/list
