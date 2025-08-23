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
	StarID    uint //`gorm:"constraint:OnDelete:CASCADE;"`// ver: https://gorm.io/docs/constraints.html
}

type user struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Rol string 
}

type deletedToken struct {
	gorm.Model
	TokenStr string `gorm:"unique"`
}
