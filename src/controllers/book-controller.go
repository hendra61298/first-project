package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hendra61298/first-project/src/models"
	"github.com/hendra61298/first-project/src/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.BaseResponse{
		Message: "Success Get ALl Book",
		Data:    newBooks,
		Status:  http.StatusOK})
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	w.Header().Set("Content-Type", "aplication/json")
	if bookDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.BaseResponse{
			Message: "Book Not Found",
			Data:    nil,
			Status:  http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.BaseResponse{
			Message: "Success Delete Book",
			Data:    bookDetails,
			Status:  http.StatusOK})
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.BaseResponse{
		Message: "Success Create Book",
		Data:    b,
		Status:  http.StatusOK})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookExists, _ := models.GetBookById(Id)
	w.Header().Set("Content-Type", "aplication/json")
	if bookExists.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.BaseResponse{
			Message: "Book Not Found",
			Data:    nil,
			Status:  http.StatusNotFound})
	} else {
		book := models.DeleteBook(Id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.BaseResponse{
			Message: "Success Delete Book",
			Data:    book,
			Status:  http.StatusOK})
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.BaseResponse{
		Message: "Success Update Book",
		Data:    bookDetails,
		Status:  http.StatusOK})
}
