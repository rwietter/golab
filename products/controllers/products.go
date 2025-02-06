package controllers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   int
	Name string
}

type Products struct {
	products []Product
}

func addProduct(p Product, prod *Products) {
	prod.products = append(prod.products, p)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := Products{}
	products.products = []Product{
		{
			ID:   1,
			Name: "Product 1",
		},
		{
			ID:   2,
			Name: "Product 2",
		},
	}

	json.NewEncoder(w).Encode(products)
}
