package db

import "gorm.io/gorm"

type Star struct {
	gorm.Model
	ID        uint
	Name      string
	SolarMas uint
}

type Planet struct {
	gorm.Model
	ID        uint
	Name      string
	Mass      int
	IsLibable bool
	StarID    uint `gorm:"constraint:OnDelete:CASCADE;"` //Is not working, o tal vez solo funciona desde el ORM?
}