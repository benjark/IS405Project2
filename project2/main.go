package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// open file
	f, err := os.Open("names.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	// initialize slice list of students
	var strSlice []string

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		strSlice = append(strSlice, rec...)

		// printing the slice
		// fmt.Println(strSlice)

	}

	// finding the length,assigning it to a variable, and printing the variable
	numStudents := len(strSlice)
	fmt.Println(numStudents)

	// Referring to a position in the slice
	studentOne := strSlice[0]
	fmt.Println(studentOne)

	//Getting random seed
	rand.Seed(time.Now().UnixNano())

	//Shuffling the slice
	rand.Shuffle(len(strSlice), func(i, j int) {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	})

	//Printing the shuffled slice
	fmt.Println(strSlice)

}
