package datastore

import (
	"log"
	"time"

	"github.com/matt-FFFFFF/bookdata-api/loader"
)

func elapsed(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	//b.Store = &loader.BooksLiteral
	defer elapsed(time.Now(), "Loading from CSV file")
	booksList := loader.CsvReader()
	b.Store = &booksList
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}
