package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	O "github.com/MoamenKhaled/data"
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

func CreateReader(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

}

func GetReaders(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(readers)

}
