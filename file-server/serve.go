package main

import (
	"net/http"
	"os"
)

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./"))
	mux.Handle("/", fs)

	http.ListenAndServe(":"+port, mux)
}