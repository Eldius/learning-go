/*
Package handlers is where I will put the handlers
*/
package handlers

import (
	"log"
	"encoding/json"
	"html/template"
	"net/http"
	"github.com/Eldius/learning-go/webapp-go/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index is the handler for index path
func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", map[string]interface{}{
		"employees": model.ListEmployees(),
	})
}

// EmployeeList returns a json with the employee list
func EmployeeList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var response []model.EmployeeTO
		for _, e := range model.ListEmployees() {
			response = append(response, e.ToTO())
		}
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("accepts:", r.Header.Get("Accepts"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
