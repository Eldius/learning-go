package routes

import (
	"net/http"
	"github.com/Eldius/webapp-go/handlers"
)

/*
LoadRoutes loads all routes
*/
func LoadRoutes() {
	http.HandleFunc("/", handlers.Index)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
