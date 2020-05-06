package main

import (
	"net/http"
	"github.com/Eldius/learning-go/webapp-go/routes"
)

func mainOld() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
