package main

import (
	"net/http"
)

func (s *server) setupRoutes() {
	http.HandleFunc("/get_all_lots", getAll)
	http.HandleFunc("/new", lotCreate)
	http.HandleFunc("/close", lotClose)
	//s.router.GET("/get_all_lots", s.getAll())
	//s.router.POST("/lot/create", s.lotCreate())
}
