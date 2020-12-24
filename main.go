package main

import (
	"net/http"

	"github.com/gorilla/mux"

	C "./controllers"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// ############# Routes For Reader ###########
	// add a reader
	// r.HandleFunc("/reader", Handler).Methods("POST")

	// // delete a reader
	// r.HandleFunc("/reader/{id}", Handler).Methods("DELETE")

	// // get readers
	// r.HandleFunc("/readers", Handler).Methods("GET")

	// // search by id
	// r.HandleFunc("/reader/{id}", Handler).Methods("GET")

	// // search by name
	// r.HandleFunc("/reader/{name}", Handler).Methods("GET")

	// ############ End Routes ###################

	// ############ Routes For Bookstore #########

	// add a book
	// r.HandleFunc("/book", Handler).Methods("POST")

	// search by id
	// r.HandleFunc("/book/{id}", Handler).Methods("GET")

	// // search by name
	// r.HandleFunc("/book/{name}", Handler).Methods("GET")

	// // get books
	r.HandleFunc("/books", C.Handler).Methods("GET")

	// // edit book
	// r.HandleFunc("/book/{id}", Handler).Methods("PUT")

	// // sort by title or publication
	// r.HandleFunc("/book/sort-by-title", Handler).Methods("GET")
	// r.HandleFunc("/book/sort-by-publication", Handler).Methods("GET")

	// ############ End Routes ####################

	// server listens on port 8080
	http.ListenAndServe(":8080", r)
}
