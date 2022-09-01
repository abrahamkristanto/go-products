package database

import (
	"fmt"
	"log"
	"os"

	"github.com/abrahamkristanto/go-products/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var db *gorm.DB

func Connect() (err error) {
	err = godotenv.Load("/Users/abrahamkristanto/practice/go-products/.env")
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		HOST        = os.Getenv("HOST")
		PORT        = os.Getenv("PORT")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host = %v port = %v user = %v password = %v dbname = %v sslmode=disable", HOST, PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
