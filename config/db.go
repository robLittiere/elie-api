package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db
}

func ConnectTestDb() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5433/postgres"))
	if err != nil {
		panic(err)
	}
	DB = db
}
