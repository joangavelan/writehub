package handler

import "net/http"

type Server struct {}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", Home)

	return http.ListenAndServe(":3000", mux)
}

func GetServer() *Server {
	return &Server{}
}