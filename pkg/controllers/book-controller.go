package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wipdev-tech/fcc-go-bookstore/pkg/models"
	"github.com/wipdev-tech/fcc-go-bookstore/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
    book := &models.Book{}
    utils.ParseBody(r, book)

    b := book.CreateBook()
    utils.MakeResponse(w, b)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    NewBooks := models.GetAllBooks()
    res, _ := json.Marshal(NewBooks)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    id, err := utils.GetBookIdFromRequest(r)
    if err != nil {
        fmt.Println("Error while parsing bookId")
    }

    bookDetails, _ := models.GetBookById(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bookDetails)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
    updateBook := &models.Book{}
    utils.ParseBody(r, updateBook)

    id, err := utils.GetBookIdFromRequest(r)
    if err != nil {
        fmt.Println("Error while parsing bookId")
    }

    bookDetails, db := models.GetBookById(id)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }

    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }

    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    db.Save(&bookDetails)

    utils.MakeResponse(w, bookDetails)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    id, err := utils.GetBookIdFromRequest(r)
    if err != nil {
        fmt.Println("Error while parsing bookId")
    }

    book := models.DeleteBook(id)
    utils.MakeResponse(w, book)
}
