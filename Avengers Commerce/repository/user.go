package repository

import (
	"ngc11/entity"

	"gorm.io/gorm"
)

func CreateUser(user *entity.Users, db *gorm.DB) error {
	if err := db.Create(&user); err != nil {
		return err.Error
	}

	return nil
}

func GetUserByUsername(username string, db *gorm.DB) (*entity.Users, error) {
	user := entity.Users{}

	if err := db.Where("username = ?", username).First(&user); err != nil {
		return &user, err.Error
	}

	return &user, nil
}
