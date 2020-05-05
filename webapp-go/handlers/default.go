/*
Package handlers is where I will put the handlers
*/
package handlers

import (
	"html/template"
	"net/http"
	"github.com/Eldius/webapp-go/employee"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index is the handler for index path
func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", map[string]interface{}{
		"employees": employee.ListEmployees(),
	})
}
