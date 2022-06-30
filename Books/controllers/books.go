package controllers

import (
	"Basic-API/Books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author"`
}

var books = []models.Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func Data_entry() {
	models.DB.Create(&books)
}
func ListBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	//models.DB.Create(&books)
	books = append(books, book)
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	for _, a := range books {
		if a.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"id":     a.ID,
				"title":  a.Title,
				"author": a.Author,
			})
		}
	}
}

func Update(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})

}

func Delete(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})

}
