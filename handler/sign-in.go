package handler

import (
	"html/template"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"web/template/layout/root-layout.html",
		"web/template/auth/layout.html",
		"web/template/auth/sign-in.html",
	))
	
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}