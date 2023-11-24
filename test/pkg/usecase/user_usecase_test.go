package usecase

import (
	"Skyline/internal/utils"
	user_models "Skyline/pkg/models/user-models"
	user_repository "Skyline/pkg/repository/user-repository"
	user_usecase "Skyline/pkg/usecases/user-usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// CASE 1
	utils.SetDatabaseConnectionForTest("../../../internal/configs")
	userRepository := user_repository.NewUserRepository()
	//	roleRepository := role.NewRoleRepository(db)
	userService := user_usecase.NewUserUsecase(userRepository)

	password, err := utils.HashPassword(utils.RandomString(10))
	arg := user_models.UserRequest{
		FirstName:  utils.RandomString(6),
		LastName:   utils.RandomString(6),
		Email:      utils.RandomEmail(),
		Password:   password,
		RePassword: password,
	}

	user, err := userService.Create(&arg)
	if err != nil {
		panic(err)
	}

	// CASE 2
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
}
