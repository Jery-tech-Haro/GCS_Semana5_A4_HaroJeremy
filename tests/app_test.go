package tests

import (
	"testing"

	app "github.com/Jery-tech-Haro/GCS_Semana5_A4_HaroJeremy/src"
)

func TestAddAndList(t *testing.T) {
	if _, err := app.AddProduct("item", 1); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if len(app.ListProducts()) < 1 {
		t.Fatalf("se esperaba al menos 1 producto")
	}
}
