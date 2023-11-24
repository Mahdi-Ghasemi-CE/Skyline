package user_usecase

import (
	"Skyline/pkg/models/user-models"
)

type UserUsecaseInterface interface {
	Create(user *user_models.UserRequest) (*user_models.UserResponse, error)
	Update(user *user_models.UserRequest) (*user_models.UserResponse, error)
	Get(id int) (*user_models.UserResponse, error)
	Delete(id int) (bool, error)
}
