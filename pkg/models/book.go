package models

import (
	"github.com/wipdev-tech/fcc-go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
    db.Create(&b)
    return b
}

func GetAllBooks() (books []Book) {
    db.Find(&books)
    return books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) (book *Book) {
    db.Where("ID = ?", id).Delete(book)
    return book
}
