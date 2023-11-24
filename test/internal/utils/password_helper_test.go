package utils

import (
	"Skyline/internal/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordHelper(t *testing.T) {
	// CASE 1 HashPassword test
	t.Run("HashPassword", func(t *testing.T) {
		password := utils.RandomString(8)
		locallyHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		require.NoError(t, err)
		require.NotEmpty(t, locallyHashedPassword)

		HashedPassword, err := utils.HashPassword(password)
		require.NoError(t, err)

		require.NotEmpty(t, HashedPassword)
	})

	// CASE 2 CheckPassword test
	t.Run("CheckPassword", func(t *testing.T) {
		password := utils.RandomString(8)
		locallyHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		require.NoError(t, err)
		require.NotEmpty(t, locallyHashedPassword)

		HashedPassword := utils.CheckPassword(password, string(locallyHashedPassword))
		require.NoError(t, err)
		require.Empty(t, HashedPassword)

		HashedPassword = utils.CheckPassword(password, utils.RandomString(8))
		require.NotEqual(t, HashedPassword, nil)
	})
}
