package handler

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"web/template/layout/root-layout.html",
		"web/template/layout/header.html",
		"web/template/home.html",
	))
	
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}