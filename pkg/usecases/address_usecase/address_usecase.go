package address_usecase

import (
	"Skyline/pkg/models/address_models"
	"Skyline/pkg/repository/address_repository"
	"fmt"
)

type addressUsecase struct {
	addressRepository address_repository.AddressRepositoryInterface
}

func NewAddressUseCase(addressRepository address_repository.AddressRepositoryInterface) AddressUsecaseInterface {
	return &addressUsecase{
		addressRepository: addressRepository,
	}
}

func (usecase addressUsecase) GetAllCountries() ([]address_models.CountryResponse, error) {
	response, err := usecase.addressRepository.GetAllCountries()
	if err != nil {
		return nil, err
	}
	fmt.Println(response)

	countryResponse := mapCountryToCountryResponse(response)

	return countryResponse, nil
}

func mapCountryToCountryResponse(response []address_models.Country) []address_models.CountryResponse {
	var countryResponse []address_models.CountryResponse

	for i := 0; i < len(response); i++ {
		countryResponse[i].Name = response[i].Name
		countryResponse[i].Code = response[i].Code
	}
	return countryResponse
}

func (usecase addressUsecase) GetAllCitiesByCountryCode(code string) ([]address_models.CityResponse, error) {
	response, err := usecase.addressRepository.GetAllCitiesByCountryCode(code)
	if err != nil {
		return nil, err
	}

	cityResponse := mapCityToCityResponses(response)

	return cityResponse, nil
}

func mapCityToCityResponses(response []address_models.City) []address_models.CityResponse {
	var cityResponse []address_models.CityResponse

	for i := 0; i < len(response); i++ {
		cityResponse[i].Name = response[i].Name
		cityResponse[i].Code = response[i].Code
	}
	return cityResponse
}

func (usecase addressUsecase) GetCity(code string) (*address_models.CityResponse, error) {
	response, err := usecase.addressRepository.GetCity(code)
	if err != nil {
		return nil, err
	}

	cityResponse := &address_models.CityResponse{
		Name: response.Name,
		Code: response.Code,
	}
	return cityResponse, nil
}

func (usecase addressUsecase) GetCountry(code string) (*address_models.CountryResponse, error) {
	response, err := usecase.addressRepository.GetCountry(code)
	if err != nil {
		return nil, err
	}

	countryResponse := &address_models.CountryResponse{
		Name: response.Name,
		Code: response.Code,
	}

	return countryResponse, nil
}
