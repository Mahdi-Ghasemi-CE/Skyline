package address_models

type City struct {
	CityId int `gorm:"primaryKey"`
	Name   string
	Code   string `gorm:"foreignKey:Code"`
}
