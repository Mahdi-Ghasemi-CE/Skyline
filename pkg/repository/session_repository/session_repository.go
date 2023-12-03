package session_repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/session_models"
	"fmt"
	"gorm.io/gorm"
)

type sessionRepository struct {
	database *gorm.DB
}

func NewSessionRepository() SessionRepositoryInterface {
	return &sessionRepository{
		database: utils.DB,
	}
}

func (repository sessionRepository) Create(session *session_models.Session) (*session_models.Session, error) {
	if err := repository.database.Create(&session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (repository sessionRepository) Update(session *session_models.Session) (*session_models.Session, error) {
	if err := repository.database.Save(&session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (repository sessionRepository) Get(id int) (*session_models.Session, error) {
	var session session_models.Session
	if err := repository.database.Where("session_id = ?", id).Find(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (repository sessionRepository) IsExist(userId int) (bool, error) {
	var session session_models.Session
	if err :=
		repository.
			database.
			Where("user_id = ?", userId).
			Where("is_blocked = ?", false).
			Find(&session).Error; err != nil {
		fmt.Println(session)
		return false, err
	}
	if session.SessionId > 0 {
		return true, error(nil)
	}
	return false, nil
}

func (repository sessionRepository) Delete(id int) (bool, error) {
	if err := repository.database.Delete(&session_models.Session{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
