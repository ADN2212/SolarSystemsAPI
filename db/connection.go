package db

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// esto deberia ser una variable de entorno.
const dsn string = "host=localhost user=postgres password=123456 dbname=SolarSystemsDB2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

var dbContext = context.Background()

var db *gorm.DB = (func() *gorm.DB {
	fmt.Println("Connecting to the database ...")

	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(".env file not found")
	}

	dataBaseDNS := os.Getenv("DATA_BASE_DNS")

	if len(dataBaseDNS) == 0 {
		panic("Data Base DNS not found")
	}

	db, err := gorm.Open(postgres.Open(dataBaseDNS), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Star{})
	db.AutoMigrate(&Planet{})

	fmt.Println("Successful connection to the database")
	return db

})()
