package models

import (
	"fmt"
	"github.com/hendra61298/first-project/src/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	gorm.Model
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("error migrate")
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).First(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
