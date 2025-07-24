package bank

type Bank struct {
	Clients      map[int]*clients.Client
	Transactions []transactions.Transaction
}
