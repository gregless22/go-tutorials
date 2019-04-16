package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	pf = fmt.Printf
	pl = fmt.Println
)

// Init Books var as a slice of Book struct

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all of the params

	// Loop through books and find with id
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all of the params

	for index, book := range books {
		if book.ID == params["id"] {
			//delete the book
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
		json.NewEncoder(w).Encode(books)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all of the params
	// Loop through books and find with id
	for index, book := range books {
		if book.ID == params["id"] {
			//delete the book
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}

//Book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	// Init the mux Router
	r := mux.NewRouter()

	//Mock Data TODO implement databae
	books = append(books, Book{ID: "1", Isbn: "1213214", Title: "Book 1", Author: &Author{FirstName: "Greg", LastName: "Connolly"}})
	books = append(books, Book{ID: "2", Isbn: "2341235", Title: "Book 2", Author: &Author{FirstName: "John", LastName: "Frank"}})
	books = append(books, Book{ID: "3", Isbn: "5346523", Title: "Book 3", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	// Route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8989", r))
}
