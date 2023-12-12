package user_usecase

import (
	"Skyline/pkg/models/user-models"
)

type UserUsecaseInterface interface {
	Create(user *user_models.UserRequest) (*user_models.UserResponse, error)
	Update(user *user_models.UpdateUserRequest) (*user_models.UserResponse, error)
	Get(id int) (*user_models.UserResponse, error)
	GetByEmail(email string) (*user_models.UserResponse, error)
	ActivateUser(email string, verifyCode int) (bool, error)
	Delete(id int) (bool, error)
}
