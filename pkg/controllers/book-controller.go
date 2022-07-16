package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/AyariAhmed/bookstore-crud/pkg/models"
	"github.com/AyariAhmed/bookstore-crud/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	retrievedBook := models.GetAllBooks()
	res, _ := json.Marshal(retrievedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		fmt.Println("error setting the response")
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	retrievedBook, _ := models.GetBookById(ID)
	res, _ := json.Marshal(retrievedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId...")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId...")
	}
	book, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
