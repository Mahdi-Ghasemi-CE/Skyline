package repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/role_models"
	"Skyline/pkg/repository/role_repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCRUDRole(t *testing.T) {
	utils.SetDatabaseConnectionForTest("../../../internal/configs")

	// CASE 0 initialize
	roleRepository := role_repository.NewRoleRepository()
	arg := &role_models.Role{
		RoleId:       utils.RandomInt(2, 1000),
		Title:        utils.RandomString(5),
		PersianTitle: utils.RandomString(5),
	}

	// CASE 1 Create
	t.Run("Create", func(t *testing.T) {
		role, err := roleRepository.Create(arg)
		if err != nil {
			panic(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, role)

		require.Equal(t, arg.RoleId, role.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 2 Get
	t.Run("Get", func(t *testing.T) {
		role, err := roleRepository.Get(arg.RoleId)
		if err != nil {
			panic(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, role)

		require.Equal(t, arg.RoleId, role.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 3 GetByTitle
	t.Run("GetByTitle", func(t *testing.T) {
		role, err := roleRepository.GetByTitle(arg.Title)
		if err != nil {
			panic(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, role)

		require.Equal(t, arg.RoleId, role.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 4 Create
	t.Run("Delete", func(t *testing.T) {
		response, err := roleRepository.Delete(arg.RoleId)
		if err != nil {
			panic(err)
		}
		role, err := roleRepository.GetByTitle(arg.Title)
		if err != nil {
			panic(err)
		}

		require.NoError(t, err)
		require.Empty(t, role)
		require.True(t, response, true)
	})
}
