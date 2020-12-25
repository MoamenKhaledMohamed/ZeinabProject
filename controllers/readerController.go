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

// my reader'sStruct
type reader struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Employment string `json:"employment"`
}

func ConvertReaderToRow(oneReader reader) string {
	row := strconv.Itoa(oneReader.Id) + " "
	row += oneReader.Name + " "
	row += oneReader.Birthday + " "
	row += strconv.Itoa(oneReader.Height) + " "
	row += strconv.Itoa(oneReader.Weight) + " "
	row += oneReader.Employment
	row += "\n"
	return row
}
func CreateReader(w http.ResponseWriter, r *http.Request) {
	// take data from body of request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// store data into oneReader
	var oneReader reader
	json.Unmarshal(reqBody, &oneReader)

	// convert data to string (row)
	row := ConvertReaderToRow(oneReader)
	// write in file
	O.AddRow("data/readers.txt", row)

	json.NewEncoder(w).Encode(oneReader)

}

func GetArrayOfReaders() []reader {
	// open file and return rows
	rows := O.GetRows("data/readers.txt")

	// array of reader'sStruct to sotre reader'sdata
	readers := []reader{}

	for _, row := range rows {
		// convert row to array of words
		words := strings.Split(row, " ")

		// cast to int
		id, _ := strconv.Atoi(words[0])
		height, _ := strconv.Atoi(words[3])
		weight, _ := strconv.Atoi(words[4])

		// create one reader then store in array of reader'sStruct
		newReader := reader{Id: id,
			Name:       words[1],
			Birthday:   words[2],
			Height:     height,
			Weight:     weight,
			Employment: words[5],
		}

		readers = append(readers, newReader)

	}
	return readers
}

func GetReaders(w http.ResponseWriter, r *http.Request) {
	readers := GetArrayOfReaders()
	json.NewEncoder(w).Encode(readers)

}

func DeleteReader(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// get all readers then search for a specific reader
	// delete the file then write again without the deleted reader
	readers := GetArrayOfReaders()
	O.RemoveFile("data/readers.txt")
	for _, reader := range readers {
		if reader.Id != id {
			row := ConvertReaderToRow(reader)
			O.AddRow("data/readers.txt", row)
		}
	}

}

func SearchById(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	readers := GetArrayOfReaders()
	for _, reader := range readers {
		if reader.Id == id {
			json.NewEncoder(w).Encode(reader)
		}
	}

}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	name := vars["name"]

	readers := GetArrayOfReaders()
	for _, reader := range readers {
		if reader.Name == name {
			json.NewEncoder(w).Encode(reader)
		}
	}
}
