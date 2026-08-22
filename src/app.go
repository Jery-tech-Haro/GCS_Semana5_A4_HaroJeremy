package app

import "errors"

type Product struct {
	Name string
	Qty  int
}

var products []Product

func ListProducts() []Product {
	return products
}

func AddProduct(name string, qty int) (bool, error) {
	if name == "" {
		return false, errors.New("name required")
	}
	if qty < 0 {
		return false, errors.New("qty must be >= 0")
	}
	products = append(products, Product{Name: name, Qty: qty})
	return true, nil
}
