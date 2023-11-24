package usecases

import (
	"Skyline/internal/customErrors"
	"Skyline/internal/utils"
	"Skyline/pkg/models"
	user_repository "Skyline/pkg/repository/user-repository"
	"gorm.io/gorm"
	"time"
)

type userUsecase struct {
	userRepository user_repository.UserRepositoryInterface
	database       *gorm.DB
}

func NewUserUsecase(userRepository user_repository.UserRepositoryInterface) UserUsecaseInterface {
	return &userUsecase{
		userRepository: userRepository,
		database:       utils.DB,
	}
}

func (service userUsecase) Create(userRequest *models.UserRequest) (*models.UserResponse, error) {
	var err error

	isDuplicate, err := service.userRepository.IsExist(userRequest.Email)
	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, customErrors.DuplicateDataError("email")
	}

	userRequest.Password, err = utils.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	/*role, err := service.roleRepository.GetByTitle("Customer")
	if err != nil {
		return nil, err
	}*/

	user := &models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		//RoleId:    role.RoleId,
		CreatedAt: time.Now(),
	}
	response, _ := service.userRepository.Create(user)
	userResponse := &models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		//RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (service userUsecase) Update(userRequest *models.UserRequest) (*models.UserResponse, error) {
	var err error

	isDuplicate, err := service.userRepository.IsExist(userRequest.Email)
	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, customErrors.DuplicateDataError("email")
	}

	userRequest.Password, err = utils.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
	}
	response, _ := service.userRepository.Update(user)
	userResponse := &models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		//RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (service userUsecase) Get(id int) (*models.UserResponse, error) {
	response, err := service.userRepository.Get(id)
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		//RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (service userUsecase) Delete(id int) (bool, error) {
	response, err := service.userRepository.Delete(id)
	if err != nil {
		return response, err
	}
	return response, err
}
