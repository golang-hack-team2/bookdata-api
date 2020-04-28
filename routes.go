package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookdata-api/loader"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}
	data := books.GetAllBooks(limit, skip)
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("skip")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
}

func getBookByISBN(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	var err error

	ISBN, ok := pathParams["isbn"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "ISBN needs to be included in the URI path"}`))
		return
	}

	data, err := books.GetBookByISBN(ISBN)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "no match for ISBN"}`))
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	var err error
	author, ok := pathParams["author"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "author string needs to be included in the URI path"}`))
		return
	}

	data, err := books.GetBooksByAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "no match for author"}`))
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getBooksByTitle(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	var err error
	title, ok := pathParams["title"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "title string needs to be included in the URI path"}`))
		return
	}

	data, err := books.GetBooksByTitle(title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "no match for title"}`))
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func createBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error reading request body"}`))
		return
	}
	var b loader.BookData
	if err = json.Unmarshal(body, &b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	books.AddBook(b)
	w.WriteHeader(http.StatusOK)
	return
}

func deleteBookByISBN(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	ISBN, ok := pathParams["isbn"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "ISBN needs to be included in the URI path"}`))
		return
	}

	_, err := books.DeleteBook(ISBN)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not find book by ISBN"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
