package session_models

import "time"

type SessionResponse struct {
	SessionId    int       `json:"-"`
	UserId       int       `json:"-" `
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"firstName"`
	ClientIp     string    `json:"clientIp" `
	IsBlocked    bool      `json:"isBlocked" `
	ExpiresAt    time.Time `json:"ExpiresAt" `
	CreatedAt    time.Time `json:"CreatedAt" `
}
