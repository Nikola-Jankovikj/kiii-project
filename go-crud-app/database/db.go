package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-crud-app/models"
	"os"
)

var DB *gorm.DB

func GetCollection() *gorm.DB {
	return DB
}

func OpenConnection() error {
	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Book{})

	return nil
}

func CloseConnection() {
	err := DB.Close()
	if err != nil {
		return
	}
}
