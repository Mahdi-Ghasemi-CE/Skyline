package user_usecase

import (
	"Skyline/internal/customErrors"
	"Skyline/internal/utils"
	"Skyline/pkg/models/user-models"
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

func (service userUsecase) Create(userRequest *user_models.UserRequest) (*user_models.UserResponse, error) {
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

	user := &user_models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		//RoleId:    role.RoleId,
		CreatedAt: time.Now(),
	}
	response, _ := service.userRepository.Create(user)
	userResponse := &user_models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		//RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (service userUsecase) Update(userRequest *user_models.UserRequest) (*user_models.UserResponse, error) {
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
	user := &user_models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
	}
	response, _ := service.userRepository.Update(user)
	userResponse := &user_models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		//RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (service userUsecase) Get(id int) (*user_models.UserResponse, error) {
	response, err := service.userRepository.Get(id)
	if err != nil {
		return nil, err
	}

	userResponse := &user_models.UserResponse{
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
