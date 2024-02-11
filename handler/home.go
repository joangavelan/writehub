package handler

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"web/template/layout.html",
		"web/template/home.html",
	))
	
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}