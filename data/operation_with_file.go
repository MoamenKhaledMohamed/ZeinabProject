package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreatFile(nameOfFile string) {

}

func RemoveFile(nameOfFile string) {
	e := os.Remove(nameOfFile)
	if e != nil {
		log.Fatal(e)
	}
}

func OpenFileByOsToRead(nameOfFile string) *os.File {
	file2, err := os.Open(nameOfFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file2
}

func OpenFileByOsToWrite(nameOfFile string) *os.File {
	file2, err := os.OpenFile(nameOfFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file2
}

func GetRows(nameOfFile string) []string {

	file := OpenFileByOsToRead(nameOfFile)

	var rows []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	defer file.Close()

	return rows

}

func AddRow(nameOfFile string, row string) {
	file := OpenFileByOsToWrite(nameOfFile)
	_, err2 := file.WriteString(row)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("done")

	}
	defer file.Close()
}
