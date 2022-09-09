package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {

// 	f, err := os.Open("names.csv")

// 	if err != nil {

// 		log.Fatal(err)
// 	}

// 	r := csv.NewReader(f)

// 	for {

// 		record, err := r.Read()

// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for value := range record {
// 			fmt.Printf("%s\n", record[value])
// 		}
// 	}
// }

// func readCsvFile(names.csv string) [][]string {
// 	f, err := os.Open(names.csv)
// 	if err != nil {
// 		log.Fatal("Unable to read input file "+names.csv, err)
// 	}
// 	defer f.Close()

// 	csvReader := csv.NewReader(f)
// 	records, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal("Unable to parse file as CSV for "+names.csv, err)
// 	}

// 	return records
// }

// func main() {
// 	records := readCsvFile("names.csv")
// 	fmt.Println(records)
// }

// func parseLocation(names.csv string) (map[string]*Point, error) {
//     f, err := os.Open(names.csv)
//     if err != nil {
//         return nil, err
//     }
//     defer f.Close()

//     csvr := csv.NewReader(f)

//     locations := map[string]*Point{}
//     for {
//         row, err := csvr.Read()
//         if err != nil {
//             if err == io.EOF {
//                 err = nil
//             }
//             return locations, err
//         }

//         p := &Point{}
//         if p.lat, err = strconv.ParseFloat(row[1], 64); err != nil {
//             return nil, err
//         }
//         if p.lon, err = strconv.ParseFloat(row[2], 64); err != nil {
//             return nil, err
//         }
//         locations[row[0]] = p
//     }
// }

func main() {
	// open file
	f, err := os.Open("names.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

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
		fmt.Printf("%+v\n", rec)
	}

}
