package repository

import (
	"ngc11/entity"

	"gorm.io/gorm"
)

func GetAllStore(db *gorm.DB) (*[]entity.Stores, error) {
	stores := []entity.Stores{}
	if err := db.Find(&stores); err != nil {
		return &stores, err.Error
	}

	return &stores, nil
}

func GetStoreById(id int, db *gorm.DB) (*entity.Stores, error) {
	store := entity.Stores{}
	if err := db.Where("id = ?", id).First(&store).Error; err != nil {
		return &store, err
	}

	return &store, nil
}
