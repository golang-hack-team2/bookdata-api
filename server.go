package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookdata-api/datastore"
)

var (
	books datastore.BookStore
)

func init() {
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()

	// Return api version
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})

	// Get the full list
	api.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)

	// Get a single book based on ISBN number
	api.HandleFunc("/book/isbn/{isbn}", getBookByISBN).Methods(http.MethodGet)

	// Case insensitive partial match on author's name
	api.HandleFunc("/books/authors/{author}", getBooksByAuthor).Methods(http.MethodGet)

	// Case insensitive partial match on titles's name
	api.HandleFunc("/books/title/{title}", getBooksByTitle).Methods(http.MethodGet)

	// Case adding a new book
	api.HandleFunc("/book", createBook).Methods(http.MethodPost)

	// Delete a single book based on ISBN number
	api.HandleFunc("/book/isbn/{isbn}", deleteBookByISBN).Methods(http.MethodDelete)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
