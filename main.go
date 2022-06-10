package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Pages  int    `json:"Pages"`
}

// Books slice
var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// add our books route and map it to our returnAllBooks function
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/book/{id}", returnSingleBook)
	myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.ID == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

}

func main() {
	Books = []Book{
		Book{Title: "The Bell Jar",
			Author: "Sylvia Plath",
			Pages:  290,
		},
		Book{Title: "Atonement",
			Author: "Ian McEwan",
			Pages:  370,
		},
		Book{Title: "Beloved",
			Author: "Toni Morrison",
			Pages:  324,
		},
	}
	handleRequests()
}
