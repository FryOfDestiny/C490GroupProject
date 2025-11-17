package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/convert", convertHandler)
	http.HandleFunc("/help", helpHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/git", gitRequestHandler)
	http.HandleFunc("/format", formatHandler)
	http.HandleFunc("/notes", notesHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handles the default case and acts as a menu
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main Page")
}

// converts a numeric input into hex,decimal,binary,
// 1's complement, 2's complement, and IEEE-754
func convertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Convert Page")
}

// Returns go help page from query string input
func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Help Page")
}

// takes a single math problem from a query string or input box and computes it
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Calculate Page")
}

// takes a git page and any number of specifications and then,
// prints specific parts of the page in a well formated manner
func gitRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Git Page")
}

// Formats the inputed go source code
func formatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Format Page")
}

// a page for creating and saving notes
// creates a file on the computer
func notesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notes Page")
}
