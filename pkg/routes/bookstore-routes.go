package routes

import (
	"github.com/MichaelGit-beep/go-bookstore/pkg/contollers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", contollers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", contollers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", contollers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", contollers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", contollers.DeleteBook).Methods("DELETE")
}
