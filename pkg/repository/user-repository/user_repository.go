package user_repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/user-models"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository() UserRepositoryInterface {
	return &userRepository{
		database: utils.DB,
	}
}

func (repository userRepository) Create(user *user_models.User) (*user_models.User, error) {
	if err := repository.database.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository userRepository) Update(user *user_models.User) (*user_models.User, error) {
	if err := repository.database.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository userRepository) Get(id int) (*user_models.User, error) {
	var user user_models.User
	if err := repository.database.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository userRepository) Delete(id int) (bool, error) {
	if err := repository.database.Delete(&user_models.User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository userRepository) IsExist(email string) (bool, error) {
	var user user_models.User
	if err := repository.database.Where("email = ?", email).Find(&user).Error; err != nil {
		return false, err
	}
	if user.UserId > 0 {
		return true, error(nil)
	}
	return false, nil
}

func (repository userRepository) GetByEmail(email string) (*user_models.User, error) {
	var user user_models.User
	if err := repository.database.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
