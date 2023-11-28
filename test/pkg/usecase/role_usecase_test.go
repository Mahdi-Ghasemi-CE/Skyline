package usecase

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/role_models"
	"Skyline/pkg/repository/role_repository"
	"Skyline/pkg/usecases/role_usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCRUDRole(t *testing.T) {
	utils.SetDatabaseConnectionForTest("../../../internal/configs")

	// CASE 0 initialize
	roleRepository := role_repository.NewRoleRepository(utils.DB)
	roleService := role_usecase.NewRoleService(roleRepository)
	arg := role_models.Role{
		RoleId:       utils.RandomInt(10, 1000),
		Title:        utils.RandomString(5),
		PersianTitle: utils.RandomString(5),
	}

	// CASE 1 create
	t.Run("Create", func(t *testing.T) {
		role, err := roleService.Create(&arg)

		require.NoError(t, err)
		require.NotEmpty(t, role)
		require.NotEmpty(t, arg.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 2 Get
	t.Run("Get", func(t *testing.T) {
		role, err, _ := roleService.Get(arg.RoleId)

		require.NoError(t, err)
		require.NotEmpty(t, role)
		require.Equal(t, arg.RoleId, arg.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 3 GetByTitle
	t.Run("GetByTitle", func(t *testing.T) {
		role, err := roleService.GetByTitle(arg.Title)

		require.NoError(t, err)
		require.NotEmpty(t, role)
		require.Equal(t, arg.RoleId, arg.RoleId)
		require.Equal(t, arg.Title, role.Title)
		require.Equal(t, arg.PersianTitle, role.PersianTitle)
	})

	// CASE 4 Delete
	t.Run("Delete", func(t *testing.T) {
		response, err := roleService.Delete(arg.RoleId)

		require.NoError(t, err)
		require.True(t, response, true)
	})
}
