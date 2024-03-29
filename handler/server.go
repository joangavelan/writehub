package handler

import "net/http"

type Server struct {}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/{$}", Home)
	mux.HandleFunc("/auth/sign-in", SignIn)
	mux.HandleFunc("/auth/sign-up", SignUp)

	return http.ListenAndServe(":3000", mux)
}

func GetServer() *Server {
	return &Server{}
}