package main

import (

	
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sravani/go-bookstore/pkg/routes"
)
func main(){
	
	routes.RegisterBookStoreRoutes()
	//http.Handle("/", r)

	

}

