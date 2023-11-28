package role_models

type Role struct {
	RoleId       int    `gorm:"primaryKey; autoIncrement:false"`
	Title        string `gorm:"uniqueIndex"`
	PersianTitle string
}
