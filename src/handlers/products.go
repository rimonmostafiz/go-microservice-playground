package handlers

import (
	"encoding/json"
	"github.com/rimonmostafiz/go-microservice-playgroud/src/data"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	listOfProducts := data.GetProducts()
	d, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	_, err = rw.Write(d)
	if err != nil {
		http.Error(rw, "Unable to write json to response writer", http.StatusInternalServerError)
	}
}
