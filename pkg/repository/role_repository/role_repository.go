package role_repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/role_models"
	"gorm.io/gorm"
)

type roleRepository struct {
	database *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepositoryInterface {
	return &roleRepository{
		database: utils.DB,
	}
}

func (repository roleRepository) Get(id int) (*role_models.Role, error) {
	var role role_models.Role
	if err := repository.database.Where("Role_Id = ?", id).Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (repository roleRepository) GetByTitle(title string) (*role_models.Role, error) {
	var role role_models.Role
	if err := repository.database.Where("Title = ?", title).Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (repository roleRepository) Create(role *role_models.Role) (*role_models.Role, error) {
	if err := repository.database.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (repository roleRepository) Delete(id int) (bool, error) {
	if err := repository.database.Delete(&role_models.Role{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
