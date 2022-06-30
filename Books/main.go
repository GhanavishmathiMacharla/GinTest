package main

import (
	"Basic-API/Books/controllers"
	"Basic-API/Books/models"

	_ "github.com/bmizerany/pq"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	models.ConnectDatabase()

	controllers.Data_entry()
	r.GET("/books", controllers.ListBooksHandler)

	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.Show)
	r.PATCH("/books/:id", controllers.Update)
	r.DELETE("/books/:id", controllers.Delete)

	r.Run()
}
