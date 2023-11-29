package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	// Default row
	rows := 10

	if len(os.Args) > 1 {
		arg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		rows = arg
	}

	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	log.Printf("Generating %d rows CSV file into data.csv...", rows)

	for i := 0; i < rows; i++ {
		// Columns: external_id, name
		record := []string{
			uuid.New().String(),
			"name-" + uuid.New().String(),
		}

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	log.Println("DONE")
}
