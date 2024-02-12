package handler

import (
	"html/template"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"web/template/layout/root-layout.html",
		"web/template/auth/layout.html",
		"web/template/auth/sign-up.html",
	))
	
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}