package usecase

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/session_models"
	user_models "Skyline/pkg/models/user-models"
	"Skyline/pkg/repository/session_repository"
	user_repository "Skyline/pkg/repository/user-repository"
	"Skyline/pkg/usecases/session_usecase"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCRUDSessionService(t *testing.T) {
	utils.SetDatabaseConnectionForTest("../../../internal/configs")

	// CASE 0 initialize
	sessionRepository := session_repository.NewSessionRepository()
	sessionUsecase := session_usecase.NewSessionUsecase(sessionRepository)

	userRepository := user_repository.NewUserRepository()
	password, err := utils.HashPassword(utils.RandomString(10))
	userArg := &user_models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomEmail(),
		Password:  password,
		CreatedAt: time.Now(),
	}
	user, err := userRepository.Create(userArg)

	sessionArg := &session_models.SessionRequest{
		UserId:       user.UserId,
		RefreshToken: utils.RandomString(15),
		UserAgent:    utils.RandomString(15),
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
	}

	// CASE 1 Create
	t.Run("Create", func(t *testing.T) {
		session, err := sessionUsecase.Create(sessionArg)
		sessionArg.SessionId = session.SessionId

		require.NoError(t, err)
		require.NotEmpty(t, session)

		require.NotEmpty(t, session.SessionId)
		require.Equal(t, sessionArg.UserId, session.UserId)
		require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
		require.Equal(t, sessionArg.UserAgent, session.UserAgent)
		require.Equal(t, sessionArg.ClientIp, session.ClientIp)
		require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)
		require.NotEmpty(t, session.ExpiresAt)
		require.NotEmpty(t, session.CreatedAt)
	})

	// CASE 2 Get
	t.Run("Get", func(t *testing.T) {
		session, err := sessionUsecase.Get(sessionArg.SessionId)
		require.NoError(t, err)
		require.NotEmpty(t, session)

		require.NotEmpty(t, session.SessionId)
		require.Equal(t, sessionArg.UserId, session.UserId)
		require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
		require.Equal(t, sessionArg.UserAgent, session.UserAgent)
		require.Equal(t, sessionArg.ClientIp, session.ClientIp)
		require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)
	})

	// CASE 3 Update
	t.Run("Update", func(t *testing.T) {
		password, err = utils.HashPassword(utils.RandomString(10))
		require.NoError(t, err)

		sessionArg.RefreshToken = utils.RandomString(15)
		sessionArg.UserAgent = utils.RandomString(15)
		sessionArg.ClientIp = "127.0.0.1"
		sessionArg.IsBlocked = false

		session, err := sessionUsecase.Update(sessionArg)
		require.NoError(t, err)
		require.NotEmpty(t, session)

		require.NotEmpty(t, session.SessionId)
		require.NotEmpty(t, session.ExpiresAt)
		require.NotEmpty(t, session.CreatedAt)
		require.Equal(t, sessionArg.UserId, session.UserId)
		require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
		require.Equal(t, sessionArg.UserAgent, session.UserAgent)
		require.Equal(t, sessionArg.ClientIp, session.ClientIp)
		require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)

		_, _ = userRepository.Delete(sessionArg.SessionId)
	})
}
