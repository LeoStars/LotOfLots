package main

import (
	"net/http"
)

func (s *server) setupRoutes() {
	http.HandleFunc("/get_all_lots", GetAll)
	http.HandleFunc("/new", lotCreate)
	http.HandleFunc("/close", lotClose)
	http.HandleFunc("/update", lotUpdate)
	http.HandleFunc("/history", getLotHistory)
}
