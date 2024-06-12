package handlers

import (
	"github.com/gin-gonic/gin"
	"go-crud-app/database"
	"go-crud-app/models"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	db := database.GetCollection()
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	db := database.GetCollection()
	db.Create(&book)
	c.JSON(http.StatusOK, book)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	db := database.GetCollection()
	db.First(&book, id)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	db := database.GetCollection()
	db.First(&book, id)
	c.BindJSON(&book)
	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	db := database.GetCollection()
	db.Delete(&book, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
