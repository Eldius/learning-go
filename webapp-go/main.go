package main

import (
	"net/http"
	"github.com/Eldius/webapp-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
