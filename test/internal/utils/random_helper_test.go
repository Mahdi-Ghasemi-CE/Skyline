package utils

import (
	"Skyline/internal/utils"
	"github.com/stretchr/testify/require"
	"reflect"
	"regexp"
	"testing"
)

func TestRandomHelper(t *testing.T) {
	// CASE 1 RandomInt test
	t.Run("RandomInt", func(t *testing.T) {
		randomIntOne := utils.RandomInt(10, 99999)
		randomIntTwo := utils.RandomInt(10, 99999)

		require.NotEmpty(t, randomIntOne)
		require.Equal(t, reflect.Int.String(), reflect.TypeOf(randomIntOne).String())
		require.NotEqual(t, randomIntOne, randomIntTwo)
		require.GreaterOrEqual(t, randomIntOne, 10)
		require.LessOrEqual(t, randomIntOne, 99999)
	})

	// CASE 2 RandomString test
	t.Run("RandomString", func(t *testing.T) {
		randomStringOne := utils.RandomString(10)
		randomStringTwo := utils.RandomString(10)

		require.NotEmpty(t, randomStringOne)
		require.Equal(t, reflect.String.String(), reflect.TypeOf(randomStringOne).String())
		require.NotEqual(t, randomStringOne, randomStringTwo)
		require.Len(t, randomStringOne, 10)
	})

	// CASE 3 RandomInt test
	t.Run("RandomInt", func(t *testing.T) {
		randomEmailOne := utils.RandomEmail()
		randomEmailTwo := utils.RandomEmail()

		require.NotEmpty(t, randomEmailOne)
		require.Equal(t, reflect.String.String(), reflect.TypeOf(randomEmailOne).String())
		require.NotEqual(t, randomEmailOne, randomEmailTwo)

		emailRegex :=
			regexp.
				MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		require.True(t, emailRegex.MatchString(randomEmailOne))
	})

	// CASE 2 RandomPersianString test
	t.Run("RandomPersianString", func(t *testing.T) {
		randomPersianStringOne := utils.RandomPersianString(10)
		randomPersianStringTwo := utils.RandomPersianString(10)

		require.NotEmpty(t, randomPersianStringOne)
		require.Equal(t, reflect.String.String(), reflect.TypeOf(randomPersianStringOne).String())
		require.NotEqual(t, randomPersianStringOne, randomPersianStringTwo)
		require.Len(t, randomPersianStringOne, 10)
	})
}
