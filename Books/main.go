package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookhandler)
	r.DELETE("/books/:id", deleteBookHandler)
	r.GET("/books/:id", show)

	r.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBookhandler(c *gin.Context) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func show(c *gin.Context) {
	id := c.Param("id")
	//var book_1 Book
	for _, a := range books {
		if a.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"id":     a.ID,
				"title":  a.Title,
				"author": a.Author,
			})
			//	book_1 = books[:i]
		}
	}
}
