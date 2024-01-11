package address_repository

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/address_models"
	"fmt"
	"gorm.io/gorm"
)

type addressRepository struct {
	database *gorm.DB
}

func NewAddressRepository() AddressRepositoryInterface {
	return &addressRepository{
		database: utils.DB,
	}
}

func (repository addressRepository) GetAllCountries() ([]address_models.Country, error) {
	var countries []address_models.Country
	if err := repository.database.Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

func (repository addressRepository) GetAllCitiesByCountryCode(code string) ([]address_models.City, error) {
	var cities []address_models.City
	query := repository.
		database.
		Joins("inner join countries on countries.Code = cities.code").
		Find(&cities).Statement.SQL.String()
	fmt.Print(query)
	if err := repository.
		database.
		Joins("inner join countries on countries.Code = cities.code").
		Find(&cities).Error; err != nil {
		//	if err := repository.database.Joins("left join emails on emails.user_id = users.id").Where("country_id = ?", countryId).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (repository addressRepository) GetCity(code string) (address_models.City, error) {
	var city address_models.City
	if err := repository.
		database.
		Where("code = ?", code).
		Find(&city).Error; err != nil {
		return city, err
	}
	return city, nil
}

func (repository addressRepository) GetCountry(code string) (address_models.Country, error) {
	var country address_models.Country
	if err := repository.
		database.
		Where("code = ?", code).
		Find(&country).Error; err != nil {
		return country, err
	}
	return country, nil
}
