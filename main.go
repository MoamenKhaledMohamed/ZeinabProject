package main

import (
	"net/http"

	"github.com/gorilla/mux"

	C "github.com/MoamenKhaled/Controllers"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// ############# Routes For Reader ###########
	// add a reader
	r.HandleFunc("/reader", C.CreateReader).Methods("POST")

	// delete a reader
	r.HandleFunc("/reader/{id}", C.DeleteReader).Methods("DELETE")

	// get readers
	r.HandleFunc("/readers", C.GetReaders).Methods("GET")

	// search by id
	r.HandleFunc("/reader/{id}", C.SearchById).Methods("GET")

	// search by name
	r.HandleFunc("/readerName/{name}", C.SearchByName).Methods("GET")

	// ############ End Routes ###################

	// ############ Routes For Bookstore #########

	// add a book
	r.HandleFunc("/book", C.CreateBook).Methods("POST")

	// search by id
	r.HandleFunc("/book/{id}", C.SearchByIdForBook).Methods("GET")

	// search by name
	r.HandleFunc("/bookTitle/{title}", C.SearchByNameForBook).Methods("GET")

	// get books
	r.HandleFunc("/books", C.GetBooks).Methods("GET")

	// edit book
	r.HandleFunc("/bookEdit/{id}", C.EditBook).Methods("PUT")

	// sort by title or publication
	r.HandleFunc("/book-sort-by-title", C.SortByTitle).Methods("GET")
	r.HandleFunc("/book-sort-by-publication", C.SortByPublication).Methods("GET")

	// ############ End Routes ####################

	// server listens on port 8080
	http.ListenAndServe(":8080", r)
}
