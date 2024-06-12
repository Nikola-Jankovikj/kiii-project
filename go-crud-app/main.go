package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-crud-app/database"
	"go-crud-app/handlers"
)

func initApp() error {

	err := loadEnv()
	if err != nil {
		return err
	}

	err = database.OpenConnection()

	if err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initApp()
	if err != nil {
		panic("failed to connect database")
	}
	defer database.CloseConnection()

	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.Run()
}
