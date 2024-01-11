package address_models

type Continent struct {
	ContinentId   int `gorm:"primaryKey"`
	Name          string
	ContinentCode string
}
