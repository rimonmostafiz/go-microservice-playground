package data

import "time"

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
