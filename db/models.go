package DB

import "gorm.io/gorm"

type star struct {
	gorm.Model
	ID        uint
	Name      string
	SolarMas uint
}

type planet struct {
	gorm.Model
	ID        uint
	Name      string
	Mass      int
	IsLibable bool
	StarID    uint `gorm:"constraint:OnDelete:CASCADE;"`//Chekea esto mas detenidamente en la Doc.
}

type user struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

//En esta tabla se guardan los tokens cuando un user hace logout
type deletedToken struct {
	gorm.Model
	TokenStr string `gorm:"unique"`
}
