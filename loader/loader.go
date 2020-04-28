package loader

// BookData is the record structure of the books datastore
type BookData struct {
	BookID        int     `json:"book_id"`
	Title         string  `json:"title"`
	Authors       string  `json:"authors"`
	AverageRating float64 `json:"average_rating"`
	ISBN          string  `json:"isbn"`
	ISBN13        string  `json:"isbn_13"`
	LanguageCode  string  `json:"language_code"`
	NumPages      int     `json:"num_pages"`
	Ratings       int     `json:"ratings"`
	Reviews       int     `json:"reviews"`
}
