package main

import (
	"log"
	"net/http"
)

type server struct {
	data     *database
}

var s = &server{}

func main() {
	s.setupRoutes()
	s.setupDB()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
