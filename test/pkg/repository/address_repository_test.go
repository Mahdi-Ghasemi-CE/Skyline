package repository

import (
	"Skyline/pkg/repository/address_repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddressRepository(t *testing.T) {
	// CASE 0 initialize
	addressRepository := address_repository.NewAddressRepository()

	// CASE 1 GetAllCountries
	t.Run("GetAllCountries", func(t *testing.T) {
		response, err := addressRepository.GetAllCountries()
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotZero(t, len(response))
		require.NotNil(t, response[2])
		require.NotNil(t, response[2].Name)
		require.NotNil(t, response[2].Code)
	})

	// CASE 2 GetCountry
	t.Run("GetCountry", func(t *testing.T) {
		response, err := addressRepository.GetCountry("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotNil(t, response)
		require.NotNil(t, response.Name)
		require.NotNil(t, response.Code)
	})

	// CASE 3 GetCity
	t.Run("GetCity", func(t *testing.T) {
		response, err := addressRepository.GetCity("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotNil(t, response)
		require.NotNil(t, response.Name)
		require.NotNil(t, response.Code)
	})

	// CASE 4 GetAllCitiesByCountryId
	t.Run("GetAllCitiesByCountryId", func(t *testing.T) {
		response, err := addressRepository.GetAllCitiesByCountryCode("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotZero(t, len(response))
		require.NotNil(t, response[2])
		require.NotNil(t, response[2].Name)
		require.NotNil(t, response[2].Code)
	})
}
