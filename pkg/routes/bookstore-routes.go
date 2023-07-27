package routes
import (
	"github.com/gorilla/mux"
	"github.com/sravani/go-bookstore/pkg/controllers"
	"fmt"
	"log"
	"net/http"
)
func RegisterBookStoreRoutes(){
	r := mux.NewRouter()
	fmt.Println("In Routes")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
    r.HandleFunc("/book/{id}",controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}",controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}",controllers.DeleteBook).Methods("DELETE")
	fmt.Println("Starting server 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}