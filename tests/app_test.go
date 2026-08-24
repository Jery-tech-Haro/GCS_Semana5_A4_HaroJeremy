package tests

import (
	"testing"

	app "github.com/Jery-tech-Haro/GCS_Semana5_A4_HaroJeremy/src"
)

func TestAddAndList(t *testing.T) {
	if _, err := app.AddProduct("item", 1, ""); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if len(app.ListProducts()) < 1 {
		t.Fatalf("se esperaba al menos 1 producto")
	}
}

func TestAddProductInvalidDate(t *testing.T) {
	if _, err := app.AddProduct("bad-date-item", 1, "not-a-date"); err == nil {
		t.Fatalf("se esperaba error por fecha invalida")
	}
}

// Criterio 1 (REQ-003): un rango desde/hasta retorna solo los productos en rango.
func TestFilterByDateRange(t *testing.T) {
	if _, err := app.AddProduct("filtro-dentro", 5, "2030-06-15"); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if _, err := app.AddProduct("filtro-fuera", 5, "2031-01-01"); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	result, err := app.FilterByDate("2030-01-01", "2030-12-31")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	found, excluded := false, false
	for _, p := range result {
		if p.Name == "filtro-dentro" {
			found = true
		}
		if p.Name == "filtro-fuera" {
			excluded = true
		}
	}
	if !found {
		t.Fatalf("se esperaba que 'filtro-dentro' estuviera en el resultado")
	}
	if excluded {
		t.Fatalf("'filtro-fuera' no debia estar en el resultado")
	}
}

// Criterio 2 (REQ-003): sin criterio de fecha, se retorna el listado completo.
func TestFilterByDateEmptyReturnsAll(t *testing.T) {
	before := len(app.ListProducts())
	result, err := app.FilterByDate("", "")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if len(result) != before {
		t.Fatalf("se esperaba el listado completo (%d), se obtuvo %d", before, len(result))
	}
}

// Criterio 3 (REQ-003): fecha invalida o desde > hasta retorna error controlado.
func TestFilterByDateInvalidRange(t *testing.T) {
	if _, err := app.FilterByDate("2030-12-31", "2030-01-01"); err == nil {
		t.Fatalf("se esperaba error por desde > hasta")
	}
	if _, err := app.FilterByDate("fecha-invalida", ""); err == nil {
		t.Fatalf("se esperaba error por formato de fecha invalido")
	}
}
