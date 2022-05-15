package models

import (
	"github.com/kuldeep/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", bookId).Find(&book)
	return &book, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID=?", bookId).Delete(book)
	return book
}