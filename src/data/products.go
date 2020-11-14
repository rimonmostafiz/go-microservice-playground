package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	createTime  string
	createBy    string
	editTime    string
	editedBy    string
	version     string
}

func GetProducts() []*Product {
	return productList
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
