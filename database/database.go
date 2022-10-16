package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbName,
	dbPassword,
	dbUser,
	dbHost,
	dbPort = os.Getenv("DATABASE_NAME"),
	os.Getenv("DATABASE_PASSWORD"),
	os.Getenv("DATABASE_USER"),
	os.Getenv("DATABASE_HOST"),
	os.Getenv("DATABASE_PORT")

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Belgrade",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to DB")
	}

	DB = db
	return DB
}
