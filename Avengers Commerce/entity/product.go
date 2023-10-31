package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID           uint
	Name         string
	Stock        int
	Price        int
	Transactions []Transactions // Has many association
}
