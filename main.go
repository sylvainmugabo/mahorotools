package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {

	if len(os.Args) < 1 {
		log.Fatal("No sufficient arguments were supplied")
	}

	fileName := os.Args[1]

	path := "finalvoters.csv"
	filePath, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File can not be open", err)
	}

	from, err := csv.NewReader(filePath).ReadAll()
	if err != nil {
		log.Fatal("Unable to read a file", err)
	}
	var data [][]string

	for index, line := range from {

		if index > 0 {
			uuidWithHyphen := uuid.New()
			uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
			data = append(data, []string{line[0], line[1], uuid, line[1], "1"})
		}
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer filePath.Close()
	w := csv.NewWriter(file)
	w.WriteAll(data)
}
