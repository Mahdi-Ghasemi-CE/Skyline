package user_usecase

import (
	custom_errors "Skyline/internal/custom-errors"
	"Skyline/internal/utils"
	"Skyline/pkg/models/user-models"
	"Skyline/pkg/repository/role_repository"
	user_repository "Skyline/pkg/repository/user-repository"
	"fmt"
	"time"
)

type userUsecase struct {
	userRepository user_repository.UserRepositoryInterface
	roleRepository role_repository.RoleRepositoryInterface
}

func NewUserUsecase(userRepository user_repository.UserRepositoryInterface, roleRepository role_repository.RoleRepositoryInterface) UserUsecaseInterface {
	return &userUsecase{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (usecase userUsecase) Create(userRequest *user_models.UserRequest) (*user_models.UserResponse, error) {
	isDuplicate, err := usecase.userRepository.IsExist(userRequest.Email)
	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, custom_errors.DuplicateDataError("email")
	}

	userRequest.Password, err = utils.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	role, err := usecase.roleRepository.GetByTitle("Customer")
	if err != nil {
		return nil, err
	}

	user := &user_models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		RoleId:    role.RoleId,
		CreatedAt: time.Now(),
	}
	response, _ := usecase.userRepository.Create(user)
	userResponse := &user_models.UserResponse{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (usecase userUsecase) Update(userRequest *user_models.UpdateUserRequest) (*user_models.UserResponse, error) {
	isDuplicate, err := usecase.userRepository.IsExist(userRequest.Email)
	if err != nil {
		return nil, err
	}
	user, err := usecase.userRepository.Get(userRequest.UserId)
	if err != nil {
		return nil, err
	}
	if isDuplicate && user.Email != userRequest.Email {
		fmt.Println(custom_errors.DuplicateDataError("email"))
		return nil, custom_errors.DuplicateDataError("email")
	}

	userRequest.Password, err = utils.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	user = mapUpdateUserRequestToUser(userRequest)
	response, _ := usecase.userRepository.Update(user)
	userResponse := mapUserToUserResponse(response)

	return userResponse, nil
}

func mapUserToUserResponse(response *user_models.User) *user_models.UserResponse {
	userResponse := &user_models.UserResponse{
		UserId:    response.UserId,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		RoleId:    response.RoleId,
	}
	return userResponse
}

func mapUpdateUserRequestToUser(updateUserRequest *user_models.UpdateUserRequest) *user_models.User {
	user := &user_models.User{
		UserId:    updateUserRequest.UserId,
		FirstName: updateUserRequest.FirstName,
		LastName:  updateUserRequest.LastName,
		Email:     updateUserRequest.Email,
		Password:  updateUserRequest.Password,
	}
	return user
}

func (usecase userUsecase) Get(id int) (*user_models.UserResponse, error) {
	response, err := usecase.userRepository.Get(id)
	if err != nil {
		return nil, err
	}
	if response.UserId == 0 {
		return &user_models.UserResponse{}, nil
	}
	userResponse := &user_models.UserResponse{
		UserId:    response.UserId,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
		RoleId:    response.RoleId,
	}

	return userResponse, nil
}

func (usecase userUsecase) Delete(id int) (bool, error) {
	response, err := usecase.userRepository.Delete(id)
	if err != nil {
		return response, err
	}
	return response, err
}
