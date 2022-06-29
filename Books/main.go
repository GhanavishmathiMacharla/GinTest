package main

import (
	//"database/sql"
	"Basic-API/Books/controllers"
	"Basic-API/Books/models"

	//"log"

	_ "github.com/bmizerany/pq"
	"github.com/gin-gonic/gin"
)

func main() {

	// host := "localhost"
	// port := "5432"
	// user := "gmacharl"
	// password := "Munny@123"
	// dbname := "book_store"
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)

	// if err != nil {
	// 	log.Fatal("Can't connect to the Database", err)
	// }
	// defer db.Close()
	// //	err = result.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully Connected")

	// sqlStatement := `INSERT INTO book_db(id,title,author)VALUES(4,'THE GODFATHER','MARIO PUZO')`
	// _, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }
	r := gin.New()

	// r.GET("/books", listBooksHandler)
	// r.POST("/books", createBookhandler)
	// r.DELETE("/books/:id", deleteBookHandler)
	// r.GET("/books/:id", show)
	// r.PATCH("/books/:id", update)
	// // sql := "select id,nama FROM student "
	// // data, err := DB.Query(sql)
	// // if err != nil {
	// // 	saveError := fmt.Sprintf("Error Query, and %s", err)
	models.ConnectDatabase()
	// // }
	controllers.Data_entry()
	r.GET("/books", controllers.ListBooksHandler)
	//r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.Show)
	r.PATCH("/books/:id", controllers.Update)
	r.DELETE("/books/:id", controllers.Delete)

	r.Run()
}

// type Book struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// var books = []Book{
// 	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
// 	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
// 	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
// }

// func listBooksHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, books)
// }

// func createBookhandler(c *gin.Context) {
// 	var book Book
// 	err := c.ShouldBindJSON(&book)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err": err.Error(),
// 		})
// 		return
// 	}
// 	books = append(books, book)
// 	c.JSON(http.StatusCreated, book)
// }

// func deleteBookHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, a := range books {
// 		if a.ID == id {
// 			books = append(books[:i], books[i+1:]...)
// 			break
// 		}
// 	}

// 	c.Status(http.StatusNoContent)
// }

// func show(c *gin.Context) {
// 	id := c.Param("id")
// 	for _, a := range books {
// 		if a.ID == id {
// 			c.JSON(http.StatusOK, gin.H{
// 				"id":     a.ID,
// 				"title":  a.Title,
// 				"author": a.Author,
// 			})
// 		}
// 	}
// }
// func update(c *gin.Context) {
// 	id := c.Param("id")
// 	fmt.Println(id)
// 	for _, a := range books {
// 		if a.ID == id {
// 			// err := c.ShouldBindJSON(&a)
// 			// if err != nil {
// 			// 	c.JSON(http.StatusBadRequest, gin.H{
// 			// 		"err": err.Error(),
// 			// 	})
// 			// 	return
// 			// }
// 			//a.Title = "The wizard"
// 			c.JSON(http.StatusOK, gin.H{
// 				"id":     a.ID,
// 				"title":  "The Wizard",
// 				"author": a.Author,
// 			})
// 		}

// 	}
// 	fmt.Println(books)
// }
