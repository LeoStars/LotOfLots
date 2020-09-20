package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type lot struct {
	ID       int
	Name string `json:"name"`
	Price string `json:"price"`
	Status bool `json:"status"`
}

func getAll (w http.ResponseWriter, r *http.Request) {
	res, err := s.data.getAllLots()
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(res)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully got lots")
}


func lotCreate(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		Name string `json:"name"`
		Price string `json:"price"`
	}
	type Res struct {
		Error string
	}
	var req Req
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	err = s.data.newLot(req.Name, req.Price)

	if err != nil {
		panic(err)
	}

	res := &Res{
		Error: "",
	}
	fmt.Fprintf(w, res.Error)
}


func lotClose(w http.ResponseWriter, r *http.Request) {
	type Res struct {
		Error string
	}
	type Req struct {
		ID int `json:"id"`
	}
	var req Req
	decoder:= json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err!= nil {
		panic(err)
	}
	fmt.Println(req.ID)
	err = s.data.closeLot(req.ID)
	res := &Res{
		Error: "",
	}
	fmt.Fprintf(w, res.Error)
}

func lotUpdate(w http.ResponseWriter, r *http.Request) {
	type Res struct {
		Error string
	}
	type Req struct {
		ID int `json:"id"`
		newPrice string `json:"new_price"`
	}
	var req Req
	decoder:= json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err!= nil {
		panic(err)
	}
	err = s.data.updatePrice(req.ID, req.newPrice)
	res := &Res{
		Error: "",
	}
	fmt.Fprintf(w, res.Error)
}