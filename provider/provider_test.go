package provider

import (
	"database/sql"
	"testing"

	"github.com/go-chi/chi"
)

func TestUpdateURL(t *testing.T) {
	mockTest := "localhost:9090"
	UpdateURL(mockTest)
	if Provider.URL != mockTest {
		t.Error("Failed to update provider URL")
	}
}

func TestUpdateDatabase(t *testing.T) {
	mockTest, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")
	UpdateDatabase(mockTest)
	if Provider.Database != mockTest {
		t.Error("Failed to update provider database")
	}
}

func TestUpdateRouter(t *testing.T) {
	mockTest := chi.NewRouter()
	UpdateRouter(mockTest)
	if Provider.Router != mockTest {
		t.Error("Failed to update provider router")
	}
}
