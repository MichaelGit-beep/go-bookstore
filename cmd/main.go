package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MichaelGit-beep/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Stating servert at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
