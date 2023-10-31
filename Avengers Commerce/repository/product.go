package repository

import (
	"ngc11/entity"

	"gorm.io/gorm"
)

func GetAllProduct(db *gorm.DB) (*[]entity.Products, error) {
	products := []entity.Products{}
	if err := db.Find(&products).Error; err != nil {
		return &products, err
	}

	return &products, nil
}
