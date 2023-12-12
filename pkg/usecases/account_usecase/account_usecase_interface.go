package account_usecase

import (
	"Skyline/pkg/models/user-models"
	"time"
)

type AccountUsecaseInterface interface {
	Login(loginRequest *user_models.LoginRequest, ClientIp string, UserAgent string) (*user_models.LoginResponse, error)
	createAccessToken(user *user_models.User) (string, error)
	createRefreshToken(user *user_models.User) (string, time.Time, error)
	ForgetPassword(email string) (bool, error)
}
