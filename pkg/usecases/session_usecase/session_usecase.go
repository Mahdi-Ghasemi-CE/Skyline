package session_usecase

import (
	"Skyline/internal/custom-errors"
	"Skyline/pkg/models/session_models"
	"Skyline/pkg/repository/session_repository"
	"time"
)

type sessionUsecase struct {
	sessionRepository session_repository.SessionRepositoryInterface
}

func NewSessionUsecase(sessionRepository session_repository.SessionRepositoryInterface) SessionUsecaseInterface {
	return &sessionUsecase{
		sessionRepository: sessionRepository,
	}
}

func (usecase sessionUsecase) Create(sessionRequest *session_models.SessionRequest) (*session_models.SessionResponse, error) {
	var sessionResponse session_models.SessionResponse
	isDuplicate, err := usecase.sessionRepository.IsExist(sessionRequest.UserId)
	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, custom_errors.DuplicateDataError("userId")
	}

	session := &session_models.Session{
		UserId:       sessionRequest.UserId,
		RefreshToken: sessionRequest.RefreshToken,
		UserAgent:    sessionRequest.UserAgent,
		ClientIp:     sessionRequest.ClientIp,
		IsBlocked:    sessionRequest.IsBlocked,
		ExpiresAt:    time.Now().Add(time.Minute * 5),
		CreatedAt:    time.Now(),
	}
	response, err := usecase.sessionRepository.Create(session)
	if err != nil {
		return &sessionResponse, err
	}

	sessionResponse = session_models.SessionResponse{
		SessionId:    response.SessionId,
		UserId:       response.UserId,
		RefreshToken: response.RefreshToken,
		UserAgent:    response.UserAgent,
		ClientIp:     response.ClientIp,
		IsBlocked:    response.IsBlocked,
		ExpiresAt:    response.ExpiresAt,
		CreatedAt:    response.CreatedAt,
	}
	return &sessionResponse, nil
}

func (usecase sessionUsecase) Update(sessionRequest *session_models.SessionRequest) (*session_models.SessionResponse, error) {
	var sessionResponse session_models.SessionResponse
	session := &session_models.Session{
		UserId:       sessionRequest.UserId,
		RefreshToken: sessionRequest.RefreshToken,
		UserAgent:    sessionRequest.UserAgent,
		ClientIp:     sessionRequest.ClientIp,
		IsBlocked:    sessionRequest.IsBlocked,
		ExpiresAt:    time.Now().Add(time.Minute * 5),
		CreatedAt:    time.Now(),
	}
	response, err := usecase.sessionRepository.Create(session)
	if err != nil {
		return &sessionResponse, err
	}

	sessionResponse = session_models.SessionResponse{
		SessionId:    response.SessionId,
		UserId:       response.UserId,
		RefreshToken: response.RefreshToken,
		UserAgent:    response.UserAgent,
		ClientIp:     response.ClientIp,
		IsBlocked:    response.IsBlocked,
		ExpiresAt:    response.ExpiresAt,
		CreatedAt:    response.CreatedAt,
	}
	return &sessionResponse, nil
}

func (usecase sessionUsecase) Get(id int) (*session_models.SessionResponse, error) {
	response, err := usecase.sessionRepository.Get(id)
	if err != nil {
		return nil, err
	}

	sessionResponse := &session_models.SessionResponse{
		SessionId:    response.SessionId,
		UserId:       response.UserId,
		RefreshToken: response.RefreshToken,
		UserAgent:    response.UserAgent,
		ClientIp:     response.ClientIp,
		IsBlocked:    response.IsBlocked,
		ExpiresAt:    response.ExpiresAt,
		CreatedAt:    response.CreatedAt,
	}
	return sessionResponse, nil
}
