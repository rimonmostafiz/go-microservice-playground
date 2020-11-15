package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Product")
	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.logger.Printf("Prod: %#v", product)
	data.AddProduct(product)
}
