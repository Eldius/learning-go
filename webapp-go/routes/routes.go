package routes

import (
	"net/http"
	"github.com/Eldius/learning-go/webapp-go/handlers"
)

/*
LoadRoutes loads all routes
*/
func LoadRoutes() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/api/employee", handlers.EmployeeList)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
