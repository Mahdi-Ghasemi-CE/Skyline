package session_repository

import (
	"Skyline/pkg/models/session_models"
)

type SessionRepositoryInterface interface {
	Create(session *session_models.Session) (*session_models.Session, error)
	Update(session *session_models.Session) (*session_models.Session, error)
	Get(id int) (*session_models.Session, error)
	GetByUserId(userId int) (*session_models.Session, error)
	IsExist(userId int) (bool, error)
	Delete(id int) (bool, error)
}
