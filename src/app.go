package app

import (
	"errors"
	"time"
)

const dateLayout = "2006-01-02"

type Product struct {
	Name      string
	Qty       int
	CreatedAt string
}

var products []Product

func ListProducts() []Product {
	return products
}

func AddProduct(name string, qty int, createdAt string) (bool, error) {
	if name == "" {
		return false, errors.New("name required")
	}
	if qty < 0 {
		return false, errors.New("qty must be >= 0")
	}
	if createdAt != "" {
		if _, err := time.Parse(dateLayout, createdAt); err != nil {
			return false, errors.New("createdAt must be ISO-8601 (YYYY-MM-DD)")
		}
	}
	products = append(products, Product{Name: name, Qty: qty, CreatedAt: createdAt})
	return true, nil
}

// FilterByDate implementa REQ-003: filtra productos por rango de fecha de alta.
// Si desde y hasta vienen vacíos, retorna el listado completo (criterio 2).
func FilterByDate(desde, hasta string) ([]Product, error) {
	if desde == "" && hasta == "" {
		return ListProducts(), nil
	}

	var desdeT, hastaT time.Time
	var err error

	if desde != "" {
		desdeT, err = time.Parse(dateLayout, desde)
		if err != nil {
			return nil, errors.New("desde must be ISO-8601 (YYYY-MM-DD)")
		}
	}
	if hasta != "" {
		hastaT, err = time.Parse(dateLayout, hasta)
		if err != nil {
			return nil, errors.New("hasta must be ISO-8601 (YYYY-MM-DD)")
		}
	}
	if desde != "" && hasta != "" && desdeT.After(hastaT) {
		return nil, errors.New("desde must not be after hasta")
	}

	var result []Product
	for _, p := range products {
		pt, err := time.Parse(dateLayout, p.CreatedAt)
		if err != nil {
			continue
		}
		if desde != "" && pt.Before(desdeT) {
			continue
		}
		if hasta != "" && pt.After(hastaT) {
			continue
		}
		result = append(result, p)
	}
	return result, nil
}
