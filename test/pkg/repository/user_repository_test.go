package repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/user-models"
	user_repository "Skyline/pkg/repository/user-repository"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCRUDUserRepository(t *testing.T) {
	// CASE 0 initialize
	utils.SetDatabaseConnectionForTest("../../../internal/configs")
	userRepository := user_repository.NewUserRepository()
	password, err := utils.HashPassword(utils.RandomString(10))
	require.NoError(t, err)

	arg := &user_models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomEmail(),
		Password:  password,
		CreatedAt: time.Now(),
	}

	// CASE 1 Create
	t.Run("Create", func(t *testing.T) {
		user, err := userRepository.Create(arg)

		require.NoError(t, err)
		require.NotEmpty(t, user)

		require.Equal(t, arg.UserId, user.UserId)
		require.Equal(t, arg.FirstName, user.FirstName)
		require.Equal(t, arg.LastName, user.LastName)
		require.Equal(t, arg.Email, user.Email)
		require.Equal(t, arg.Password, user.Password)
		require.NotZero(t, user.CreatedAt)
	})

	// CASE 2 Get
	t.Run("Get", func(t *testing.T) {
		user, err := userRepository.Get(arg.UserId)

		require.NoError(t, err)
		require.NotEmpty(t, user)

		require.Equal(t, arg.UserId, user.UserId)
		require.Equal(t, arg.FirstName, user.FirstName)
		require.Equal(t, arg.LastName, user.LastName)
		require.Equal(t, arg.Email, user.Email)
		require.Equal(t, arg.Password, user.Password)
		require.NotZero(t, user.CreatedAt)
	})

	// CASE 3 Update
	t.Run("Update", func(t *testing.T) {
		password, err := utils.HashPassword(utils.RandomString(10))
		require.NoError(t, err)

		arg.FirstName = utils.RandomString(6)
		arg.LastName = utils.RandomString(6)
		arg.Email = utils.RandomEmail()
		arg.Password = password

		user, err := userRepository.Update(arg)
		require.NoError(t, err)
		require.NotEmpty(t, user)

		require.Equal(t, arg.UserId, user.UserId)
		require.Equal(t, arg.FirstName, user.FirstName)
		require.Equal(t, arg.LastName, user.LastName)
		require.Equal(t, arg.Email, user.Email)
		require.Equal(t, arg.Password, user.Password)
		require.NotZero(t, user.CreatedAt)
	})

	// CASE 4 Delete
	t.Run("Delete", func(t *testing.T) {
		response, err := userRepository.Delete(arg.UserId)
		require.NoError(t, err)

		user, err := userRepository.Get(arg.UserId)
		require.NoError(t, err)

		require.Empty(t, user)
		require.True(t, response, true)
	})
}

func TestIsExistUserRepository(t *testing.T) {
	// CASE 0 initialize
	userRepository := user_repository.NewUserRepository()
	password, err := utils.HashPassword(utils.RandomString(10))
	require.NoError(t, err)

	arg := &user_models.User{
		FirstName: utils.RandomString(6),
		LastName:  utils.RandomString(6),
		Email:     utils.RandomEmail(),
		Password:  password,
		CreatedAt: time.Now(),
	}

	// CASE 1 create first user
	user, err := userRepository.Create(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserId, user.UserId)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.NotZero(t, user.CreatedAt)

	// CASE 2 check the first user email in IsExist function
	isExist, err := userRepository.IsExist(arg.Email)

	require.NoError(t, err)
	require.NotEmpty(t, isExist)
	require.True(t, true)
}
