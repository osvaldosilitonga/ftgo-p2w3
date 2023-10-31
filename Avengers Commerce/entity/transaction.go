package entity

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	ID          uint
	UserID      uint
	ProductID   uint
	Quantity    int
	TotalAmount int
	User        Users    // Belongs to association
	Product     Products // Belongs to association
}
