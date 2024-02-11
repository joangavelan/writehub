package main

import (
	"log"

	"github.com/joangavelan/writehub/handler"
)

func main() {
	server := handler.GetServer()

	log.Fatal(server.Start())
}