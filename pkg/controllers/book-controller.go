package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/sravani/go-bookstore/pkg/utils"
	"github.com/sravani/go-bookstore/pkg/models"

)

var NewBook models.Book
func GetBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("in controllers get book func")
	newBooks := models.GetAllBooks()
	res, _:= json.Marshal(newBooks)
	w.Header().Set("Content - Type","pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println("in get book func")
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")

	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content - Type","pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID , err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content - Type","pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = models.Book{}
	utils.ParseBody(r, updateBook)

	fmt.Println("in update book func", updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
	fmt.Println(ID)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails , db := models.GetBookById(ID)
	fmt.Println("update book name",updateBook.Book_name)
	if updateBook.Book_name != ""{
		bookDetails.Book_name = updateBook.Book_name
		fmt.Println(updateBook.Book_name, bookDetails.Book_name)

	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content - Type","pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


