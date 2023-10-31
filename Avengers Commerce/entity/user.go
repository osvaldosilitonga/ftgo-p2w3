package entity

type Users struct {
	ID            uint
	Username      string
	Password      string `json:"-"`
	DepositAmount int
	Transactions  []Transactions `json:"-"`
}
