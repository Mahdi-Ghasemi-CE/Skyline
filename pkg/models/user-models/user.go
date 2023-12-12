package user_models

import "time"

type User struct {
	UserId     int `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string `gorm:"uniqueIndex"`
	Password   string
	IsActive   bool
	VerifyCode int
	CreatedAt  time.Time
	RoleId     int `gorm:"foreignKey:RoleId"`
}
