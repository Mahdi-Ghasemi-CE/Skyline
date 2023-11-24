package usecases

import "Skyline/pkg/models"

type UserUsecaseInterface interface {
	Create(user *models.UserRequest) (*models.UserResponse, error)
	Update(user *models.UserRequest) (*models.UserResponse, error)
	Get(id int) (*models.UserResponse, error)
	Delete(id int) (bool, error)
}
