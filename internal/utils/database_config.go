package utils

import (
	"Skyline/pkg/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(connection string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}
	AutoMigrate()
	return nil
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
}
