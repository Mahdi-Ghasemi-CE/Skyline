package address_repository

import "Skyline/pkg/models/address_models"

type AddressRepositoryInterface interface {
	GetAllCountries() ([]address_models.Country, error)
	GetAllCitiesByCountryCode(code string) ([]address_models.City, error)
	GetCity(code string) (address_models.City, error)
	GetCountry(code string) (address_models.Country, error)
}
