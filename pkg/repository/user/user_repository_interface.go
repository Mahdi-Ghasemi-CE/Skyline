package user

import "Skyline/pkg/models"

type UserRepositoryInterface interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Get(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	IsExist(email string) (bool, error)
	Delete(id int) (bool, error)
}
