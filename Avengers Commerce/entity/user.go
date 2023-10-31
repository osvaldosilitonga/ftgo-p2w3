package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID            uint
	Username      string
	Password      string
	DepositAmount int
	Transactions  []Transactions // Has many association
}
