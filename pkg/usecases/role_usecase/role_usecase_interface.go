package role_usecase

import "Skyline/pkg/models/role_models"

type RoleUsecaseInterface interface {
	Create(role *role_models.Role) (*role_models.RoleResponse, error)
	Get(id int) (*role_models.RoleResponse, error, int)
	GetByTitle(title string) (*role_models.RoleResponse, error)
	Delete(id int) (bool, error)
}
