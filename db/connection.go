package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"context"
	"fmt"
)

// esto deberia ser una variable de entorno.
const dsn string = "host=localhost user=postgres password=123456 dbname=SolarSystemsDB2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

var dbContext = context.Background()

var db *gorm.DB = (func() *gorm.DB {
	fmt.Println("Initializing database")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Star{})
	db.AutoMigrate(&Planet{})

	return db
})()
