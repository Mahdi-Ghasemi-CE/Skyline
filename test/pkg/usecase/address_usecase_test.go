package usecase

import (
	"Skyline/pkg/repository/address_repository"
	"Skyline/pkg/usecases/address_usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddressUsecase(t *testing.T) {
	// CASE 0 initialize
	addressRepository := address_repository.NewAddressRepository()
	addressUsecase := address_usecase.NewAddressUseCase(addressRepository)

	// CASE 1 GetAllCountries
	t.Run("GetAllCountries", func(t *testing.T) {
		response, err := addressUsecase.GetAllCountries()
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotZero(t, len(response))
	})

	// CASE 2 GetCountry
	t.Run("GetCountry", func(t *testing.T) {
		response, err := addressUsecase.GetCountry("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotNil(t, response)
		require.NotNil(t, response.Name)
		require.NotNil(t, response.Code)
	})

	// CASE 3 GetCity
	t.Run("GetCity", func(t *testing.T) {
		response, err := addressUsecase.GetCity("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotNil(t, response)
		require.NotNil(t, response.Name)
		require.NotNil(t, response.Code)
	})

	// CASE 4 GetAllCitiesByCountryId
	t.Run("GetAllCitiesByCountryId", func(t *testing.T) {
		response, err := addressUsecase.GetAllCitiesByCountryCode("IR")
		require.NoError(t, err)
		require.NotEmpty(t, response)

		require.NotZero(t, len(response))
		require.NotNil(t, response[2])
		require.NotNil(t, response[2].Name)
		require.NotNil(t, response[2].Code)
	})
}
