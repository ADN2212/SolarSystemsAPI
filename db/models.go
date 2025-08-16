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
	StarID    uint `gorm:"constraint:OnDelete:CASCADE;"`
}