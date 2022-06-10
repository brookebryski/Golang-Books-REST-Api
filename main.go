package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Book struct
type Book struct {
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
	http.HandleFunc("/", homePage)
	// add our books route and map it to our returnAllBooks function
	http.HandleFunc("/books", returnAllBooks)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
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
