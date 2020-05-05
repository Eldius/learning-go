package main

import (
	"html/template"
	"net/http"
	"net/url"
	"fmt"
	"os"
)

var tpl = template.Must(template.ParseFiles("template/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	params := u.Query()
	searchKey := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is: ", searchKey)
	fmt.Println("Results page is: ", page)
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// Add the following two lines
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("static/", http.StripPrefix("/static`/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler)

	http.ListenAndServe(":"+port, mux)
}
