package session_usecase

import "Skyline/pkg/models/session_models"

type SessionUsecaseInterface interface {
	Create(user *session_models.SessionRequest) (*session_models.SessionResponse, error)
	Update(user *session_models.SessionRequest) (*session_models.SessionResponse, error)
	Get(id int) (*session_models.SessionResponse, error)
}
