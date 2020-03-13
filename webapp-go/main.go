package main

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// get the book
		// navigate to the page
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"] // the page
        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
    http.ListenAndServe(":8080", nil)
}
