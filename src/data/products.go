package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	createTime  string  `json:"-"`
	createBy    string  `json:"-"`
	editTime    string  `json:"-"`
	editedBy    string  `json:"-"`
	version     string  `json:"-"`
}

type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r);
	return d.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product)  {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lastProduct := productList[len(productList) - 1]
	return lastProduct.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		createTime:  time.Now().UTC().String(),
		createBy:    "System",
		version:     "0",
	},
	&Product{
		ID:          2,
		Name:        "Expresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		createTime:  time.Now().UTC().String(),
		createBy:    "System",
		version:     "0",
	},
}
