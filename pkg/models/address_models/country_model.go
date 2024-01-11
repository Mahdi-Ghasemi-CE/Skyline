package address_models

type Country struct {
	CountryId int `gorm:"primaryKey"`
	Name      string
	Code      string
	Continent string `gorm:"foreignKey:ContinentCode"`
	Filename  string
}
