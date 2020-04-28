package datastore

import (
	"errors"
	"fmt"
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
	Store         *[]*loader.BookData `json:"store"`
	HighestBookID int
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	//b.Store = &loader.BooksLiteral
	defer elapsed(time.Now(), "Loading from CSV file")
	booksList, highestBookID := loader.CsvReader()
	b.Store = &booksList
	b.HighestBookID = highestBookID
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}

// GetBooksByAuthor returns a list of structs if it partial matches the author
func (b *Books) GetBooksByAuthor(author string) (*[]*loader.BookData, error) {
	results := make([]*loader.BookData, 0)

	for _, p := range *b.Store {
		if ToLowerContains(p.Authors, author) {
			results = append(results, p)
		}
	}

	return &results, nil
}

// GetBooksByTitle returns a list of structs if it partial matches the title
func (b *Books) GetBooksByTitle(title string) (*[]*loader.BookData, error) {
	results := make([]*loader.BookData, 0)

	for _, p := range *b.Store {
		if ToLowerContains(p.Title, title) {
			results = append(results, p)
		}
	}

	return &results, nil
}

// GetBookByISBN returns a single struct if it matches the ISBN string
func (b *Books) GetBookByISBN(isbn string) (*loader.BookData, error) {
	for _, p := range *b.Store {
		if p.ISBN == isbn {
			return p, nil
		}
	}

	// If we have got this far then we've failed to match
	return nil, errors.New("failed to match ISBN in the list of books")
}

// AddBook Adding a new book
func (b *Books) AddBook(book loader.BookData) (*loader.BookData, error) {

	_, err := b.GetBookByISBN(book.ISBN)
	if err == nil {
		return nil, errors.New("Book matching ISBN already exists")
	}

	newHighestBookID := b.HighestBookID + 1
	book.BookID = newHighestBookID
	b.HighestBookID = newHighestBookID

	storeWithNewBookAdded := append(*b.Store, &book)
	b.Store = &storeWithNewBookAdded

	return &book, nil
}

// DeleteBook Deleting a book
func (b *Books) DeleteBook(isbn string) (*loader.BookData, error) {
	for i, book := range *b.Store {
		if book.ISBN == isbn {
			copy((*b.Store)[i:], (*b.Store)[i+1:])
			(*b.Store)[len(*b.Store)-1] = nil
			*b.Store = (*b.Store)[:len(*b.Store)-1]
			return book, nil
		}
	}

	errmsg := fmt.Sprintf("book with '%s' ISBN not found", isbn)
	return nil, errors.New(errmsg)
}
