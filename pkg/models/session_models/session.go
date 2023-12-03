package session_models

import "time"

type Session struct {
	SessionId    int `gorm:"primaryKey"`
	UserId       int `gorm:"foreignKey:UserId"`
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
