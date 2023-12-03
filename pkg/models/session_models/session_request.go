package session_models

type SessionRequest struct {
	SessionId    int    `json:"-"`
	UserId       int    `json:"userId" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
	UserAgent    string `json:"firstName" validate:"required"`
	ClientIp     string `json:"clientIp" validate:"required"`
	IsBlocked    bool   `json:"isBlocked" validate:"required"`
}
