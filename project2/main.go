package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("names.csv")
	csvFile, _ := os.Open("names.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ','
	reader.LazyQuotes = true

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}
}
