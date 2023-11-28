package role_repository

import (
	"Skyline/pkg/models/role_models"
)

type RoleRepositoryInterface interface {
	Create(role *role_models.Role) (*role_models.Role, error)
	Get(id int) (*role_models.Role, error)
	GetByTitle(title string) (*role_models.Role, error)
	Delete(id int) (bool, error)
}
