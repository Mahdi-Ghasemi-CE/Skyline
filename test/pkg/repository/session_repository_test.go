package repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/session_models"
	user_models "Skyline/pkg/models/user-models"
	"Skyline/pkg/repository/session_repository"
	user_repository "Skyline/pkg/repository/user-repository"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSessionCRUDRepository(t *testing.T) {
	utils.SetDatabaseConnectionForTest("../../../internal/configs")

	// CASE 0 initialize
	sessionRepository := session_repository.NewSessionRepository()

	userRepository := user_repository.NewUserRepository()
	password, _ := utils.HashPassword(utils.RandomString(10))
	userArg := &user_models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomEmail(),
		Password:  password,
		CreatedAt: time.Now(),
	}

	user, err := userRepository.Create(userArg)

	sessionArg := &session_models.Session{
		UserId:       user.UserId,
		RefreshToken: utils.RandomString(15),
		UserAgent:    utils.RandomString(15),
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Duration(10)),
		CreatedAt:    time.Now(),
	}

	// CASE 1 Create
	t.Run("Create", func(t *testing.T) {
		session, err := sessionRepository.Create(sessionArg)
		require.NoError(t, err)
		require.NotEmpty(t, session)

		require.NotEmpty(t, session.SessionId)
		require.Equal(t, sessionArg.UserId, session.UserId)
		require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
		require.Equal(t, sessionArg.UserAgent, session.UserAgent)
		require.Equal(t, sessionArg.ClientIp, session.ClientIp)
		require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)
		require.Equal(t, sessionArg.ExpiresAt, session.ExpiresAt)
		require.Equal(t, sessionArg.CreatedAt, session.CreatedAt)
	})

	// CASE 2 Get
	t.Run("Get", func(t *testing.T) {
		session, err := sessionRepository.Get(sessionArg.SessionId)
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
		sessionArg.ExpiresAt = time.Now().Add(time.Duration(10))
		sessionArg.CreatedAt = time.Now()

		session, err := sessionRepository.Update(sessionArg)
		require.NoError(t, err)
		require.NotEmpty(t, session)

		require.NotEmpty(t, session.SessionId)
		require.Equal(t, sessionArg.UserId, session.UserId)
		require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
		require.Equal(t, sessionArg.UserAgent, session.UserAgent)
		require.Equal(t, sessionArg.ClientIp, session.ClientIp)
		require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)
		require.Equal(t, sessionArg.ExpiresAt, session.ExpiresAt)
	})

	// CASE 4 Delete
	t.Run("Delete", func(t *testing.T) {
		response, err := sessionRepository.Delete(sessionArg.SessionId)
		require.NoError(t, err)

		session, err := sessionRepository.Get(sessionArg.SessionId)
		require.NoError(t, err)

		require.Empty(t, session)
		require.True(t, response, true)
	})
}

func TestIsExistSessionRepository(t *testing.T) {
	utils.SetDatabaseConnectionForTest("../../../internal/configs")

	// CASE 0 initialize
	sessionRepository := session_repository.NewSessionRepository()

	userRepository := user_repository.NewUserRepository()
	password, _ := utils.HashPassword(utils.RandomString(10))
	userArg := &user_models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomEmail(),
		Password:  password,
		CreatedAt: time.Now(),
	}

	user, err := userRepository.Create(userArg)

	sessionArg := &session_models.Session{
		UserId:       user.UserId,
		RefreshToken: utils.RandomString(15),
		UserAgent:    utils.RandomString(15),
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Minute * 5),
		CreatedAt:    time.Now(),
	}

	// CASE 1 create first session
	session, err := sessionRepository.Create(sessionArg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.NotEmpty(t, session.SessionId)
	require.Equal(t, sessionArg.UserId, session.UserId)
	require.Equal(t, sessionArg.RefreshToken, session.RefreshToken)
	require.Equal(t, sessionArg.UserAgent, session.UserAgent)
	require.Equal(t, sessionArg.ClientIp, session.ClientIp)
	require.Equal(t, sessionArg.IsBlocked, session.IsBlocked)
	require.Equal(t, sessionArg.ExpiresAt, session.ExpiresAt)

	// CASE 2 check the first session userId in IsExist function
	isExist, err := sessionRepository.IsExist(sessionArg.UserId)

	require.NoError(t, err)
	require.NotEmpty(t, isExist)
	require.True(t, true)
}
