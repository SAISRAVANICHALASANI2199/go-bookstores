package models

import(
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/sravani/go-bookstore/pkg/config"
)

var db *gorm.DB


type Book struct {
	Id int `json:"id"`
	Book_name string `json:"book_name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
		db = config.GetDB()
		db.AutoMigrate(&Book{})

}
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	fmt.Println("in get book function" , Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook (ID int64) Book{
	var book Book
	db.Where("ID=?", ID). Delete(book)
	return book
}