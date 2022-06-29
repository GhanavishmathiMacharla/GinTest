package models

import (
	"fmt"
	"log"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := "localhost"
	port := "5432"
	user := "gmacharl"
	password := "Munny@123"
	dbname := "book_store"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to the Database", err)
	}
	database.AutoMigrate(&Book{})
	DB = database
}
