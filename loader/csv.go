package loader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func CsvReader() []*BookData {
	BooksCsv := make([]*BookData, 0)

	CsvFile, _ := os.Open("assets/books.csv")
	r := csv.NewReader(CsvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		avgRating, _ := strconv.ParseFloat(record[3], 64)
		numPages, _ := strconv.Atoi(record[7])
		ratings, _ := strconv.Atoi(record[8])
		reviews, _ := strconv.Atoi(record[9])

		BookRecord := BookData{
			BookID:        record[0],
			Title:         record[1],
			Authors:       record[2],
			AverageRating: avgRating,
			ISBN:          record[4],
			ISBN13:        record[5],
			LanguageCode:  record[6],
			NumPages:      numPages,
			Ratings:       ratings,
			Reviews:       reviews,
		}

		BooksCsv = append(BooksCsv, &BookRecord)

	}

	return BooksCsv
}
