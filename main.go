package main

import (
  "encoding/json"
  "log"
  "math/rand"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

// Book Struct
type Book struct {
  Id     string  `json:"id"`
  Isbn   string  `json:"isbn"`
  Title  string  `json:"title"`
  Author *Author `json:"author"`
}

// Author Struct
type Author struct {
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
}

var books []Book

// GET all books
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

// GET single book by id
func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for _, item := range books {
    if item.Id == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

//POST create book
func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.Id = strconv.Itoa(rand.Intn(1000000))
  books = append(books, book)
  json.NewEncoder(w).Encode(book)
}

//PUT update book
func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for indx, item := range books {
    if item.Id == params["id"] {
      books = append(books[:indx], books[indx+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.Id = item.Id
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  // json.NewEncoder(w).Encode(books)
}

//DELETE delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for indx, item := range books {
    if item.Id == params["id"] {
      books = append(books[:indx], books[indx+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
}

func main() {
  books = append(books, Book{Id: "1", Isbn: "8472365", Title: "Book One", Author: &Author{FirstName: "Prasad", LastName: "Surase"}})
  books = append(books, Book{Id: "2", Isbn: "2342322", Title: "Book Two", Author: &Author{FirstName: "Pratik", LastName: "More"}})

  // Init Router
  r := mux.NewRouter()

  // Define endpoints
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
