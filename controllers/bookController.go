package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	O "github.com/MoamenKhaled/data"

	"github.com/gorilla/mux"
)

// id, title, publication date, author, genre, publisher and language.
// my book'sStruct
type book struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Publicaton string `json:"publication"`
	Date       string `json:"date"`
	Author     string `json:"author"`
	Genre      string `json:"genre"`
	Publisher  string `json:"publisher"`
	Language   string `json:"language"`
}

func ConvertBookToRow(oneBook book) string {
	row := strconv.Itoa(oneBook.Id) + " "
	row += oneBook.Title + " "
	row += oneBook.Publicaton + " "
	row += oneBook.Date + " "
	row += oneBook.Author + " "
	row += oneBook.Genre + " "
	row += oneBook.Publisher + " "
	row += oneBook.Language + " "
	row += "\n"
	return row
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// take data from body of request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// store data into oneBook
	var oneBook book
	json.Unmarshal(reqBody, &oneBook)

	// convert data to string (row)
	row := ConvertBookToRow(oneBook)
	// write in file
	O.AddRow("data/books.txt", row)

	json.NewEncoder(w).Encode(oneBook)

}

func GetArrayOfBooks() []book {
	// open file and return rows
	rows := O.GetRows("data/books.txt")

	// array of reader'sStruct to sotre reader'sdata
	books := []book{}

	for _, row := range rows {
		// convert row to array of words
		words := strings.Split(row, " ")

		// cast to int
		id, _ := strconv.Atoi(words[0])

		// create one reader then store in array of reader'sStruct
		newBook := book{Id: id,
			Title:      words[1],
			Publicaton: words[2],
			Date:       words[3],
			Author:     words[4],
			Genre:      words[5],
			Publisher:  words[6],
			Language:   words[7],
		}

		books = append(books, newBook)

	}
	return books
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := GetArrayOfBooks()
	json.NewEncoder(w).Encode(books)

}

func SearchByIdForBook(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	books := GetArrayOfBooks()
	for _, book := range books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
		}
	}

}

func SearchByNameForBook(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	title := vars["title"]

	books := GetArrayOfBooks()
	for _, book := range books {
		if book.Title == title {
			json.NewEncoder(w).Encode(book)
		}
	}
}
