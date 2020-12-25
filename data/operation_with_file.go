package data

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFileByOs(nameOfFile string) *os.File {
	file2, err := os.Open(nameOfFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file2
}

func GetRows(nameOfFile string) []string {

	file := OpenFileByOs(nameOfFile)

	var rows []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	defer file.Close()

	return rows

}
