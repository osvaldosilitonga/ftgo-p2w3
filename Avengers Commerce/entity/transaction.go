package entity

type Transactions struct {
	ID      uint
	UsersID uint
	// ProductsID  uint
	// Quantity    int
	TotalAmount int
	Users       Users // Belongs to association
	// Products    Products // Belongs to association
	TransactionDetails []TransactionDetails
}

type TransactionDetails struct {
	ID             uint
	TransactionsID uint
	ProductsID     uint
	Qty            int
	Transactions   Transactions
	Products       Products
}
