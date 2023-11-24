package utils

import (
	"Skyline/pkg/models/user-models"
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
	DB.AutoMigrate(&user_models.User{})
}

func SetDatabaseConnectionForTest(path string) error {
	appConfig, err := LoadAppConfig(path)
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(postgres.Open(appConfig.DbConnection), &gorm.Config{})
	if err != nil {
		panic("Database isn't connected.")
	}
	AutoMigrate()
	return err
}
