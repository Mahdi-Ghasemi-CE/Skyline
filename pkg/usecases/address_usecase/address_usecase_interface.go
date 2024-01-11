package address_usecase

import "Skyline/pkg/models/address_models"

type AddressUsecaseInterface interface {
	GetAllCountries() ([]address_models.CountryResponse, error)
	GetAllCitiesByCountryCode(code string) ([]address_models.CityResponse, error)
	GetCity(code string) (*address_models.CityResponse, error)
	GetCountry(code string) (*address_models.CountryResponse, error)
}
