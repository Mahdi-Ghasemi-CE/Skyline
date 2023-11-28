package role_usecase

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/role_models"
	"Skyline/pkg/repository/role_repository"
	"gorm.io/gorm"
	"net/http"
)

type roleUsecase struct {
	roleRepository role_repository.RoleRepositoryInterface
	database       *gorm.DB
}

func NewRoleService(roleRepository role_repository.RoleRepositoryInterface) RoleUsecaseInterface {
	return &roleUsecase{
		roleRepository: roleRepository,
		database:       utils.DB,
	}
}

func (usecase roleUsecase) Get(id int) (*role_models.RoleResponse, error, int) {
	role, err := usecase.roleRepository.Get(id)
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	roleResponse := &role_models.RoleResponse{
		Title:        role.Title,
		PersianTitle: role.PersianTitle,
	}

	return roleResponse, nil, http.StatusOK
}

func (usecase roleUsecase) GetByTitle(title string) (*role_models.RoleResponse, error) {
	role, err := usecase.roleRepository.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	roleResponse := &role_models.RoleResponse{
		Title:        role.Title,
		PersianTitle: role.PersianTitle,
	}

	return roleResponse, nil
}

func (usecase roleUsecase) Create(role *role_models.Role) (*role_models.RoleResponse, error) {
	response, err := usecase.roleRepository.Create(role)
	if err != nil {
		return nil, err
	}
	roleResponse := &role_models.RoleResponse{
		Title:        response.Title,
		PersianTitle: response.PersianTitle,
	}
	return roleResponse, err
}

func (usecase roleUsecase) Delete(id int) (bool, error) {
	response, err := usecase.roleRepository.Delete(id)
	if err != nil {
		return response, err
	}
	return response, err
}
