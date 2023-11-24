package user_repository

import (
	"Skyline/pkg/models/user-models"
)

type UserRepositoryInterface interface {
	Create(user *user_models.User) (*user_models.User, error)
	Update(user *user_models.User) (*user_models.User, error)
	Get(id int) (*user_models.User, error)
	GetByEmail(email string) (*user_models.User, error)
	IsExist(email string) (bool, error)
	Delete(id int) (bool, error)
}
