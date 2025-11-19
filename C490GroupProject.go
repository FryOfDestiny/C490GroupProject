// Made by: Nicholas Bush, Jacob Friberg, and Logan Doutlick
//
// Description: A C490 Group Project where we create a wesbite
// to help in the creation of code and similar projects.

package main

import (
	"html/template"
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
	main, err := template.ParseFiles("main.html")
	checkError(err)

	err = main.Execute(w, nil)
	checkError(err)
}

// converts a numeric input into hex,decimal,binary,
// 1's complement, 2's complement, and IEEE-754
func convertHandler(w http.ResponseWriter, r *http.Request) {
	convert, err := template.ParseFiles("convert.html")
	checkError(err)

	err = convert.Execute(w, nil)
	checkError(err)
}

// Returns go help page from query string input
func helpHandler(w http.ResponseWriter, r *http.Request) {
	help, err := template.ParseFiles("help.html")
	checkError(err)

	err = help.Execute(w, nil)
	checkError(err)
}

// takes a single math problem from a query string or input box and computes it
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	calculate, err := template.ParseFiles("calculate.html")
	checkError(err)

	err = calculate.Execute(w, nil)
	checkError(err)
}

// takes a git page and any number of specifications and then,
// prints specific parts of the page in a well formated manner
func gitRequestHandler(w http.ResponseWriter, r *http.Request) {
	git, err := template.ParseFiles("git.html")
	checkError(err)

	err = git.Execute(w, nil)
	checkError(err)
}

// Formats the inputed go source code
func formatHandler(w http.ResponseWriter, r *http.Request) {
	format, err := template.ParseFiles("format.html")
	checkError(err)

	err = format.Execute(w, nil)
	checkError(err)
}

// a page for creating and saving notes
// creates a file on the computer
func notesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := template.ParseFiles("notes.html")
	checkError(err)

	err = notes.Execute(w, nil)
	checkError(err)
}

// function to check erros
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
