package tax

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	provider "tax-calculator/provider"

	_ "github.com/lib/pq"
)

func TestGetTotal(t *testing.T) {
	write := httptest.NewRecorder()
	read := httptest.NewRequest("GET", "/item", nil)

	mockTest, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")
	provider.UpdateDatabase(mockTest)

	GetTotal(write, read)
}

func TestPostItem(t *testing.T) {
	write := httptest.NewRecorder()
	read := httptest.NewRequest("POST", "/item", nil)

	mockTest, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")
	provider.UpdateDatabase(mockTest)

	PostItem(write, read)
}

func TestValidateNewItemOK(t *testing.T) {
	mockItem := item{
		Name:    "test",
		TaxCode: 1,
		Amount:  100,
	}
	if _, ok := validateNewItem(mockItem); !ok {
		t.Error("Failed to validate correct item")
	}
}

func TestValidateNewItemErr(t *testing.T) {
	mockItem := item{
		Name: "test",
	}
	if _, ok := validateNewItem(mockItem); ok {
		t.Error("Failed to validate incorrect item")
		// return
	}

	mockItem = item{
		TaxCode: 1,
	}
	if _, ok := validateNewItem(mockItem); ok {
		t.Error("Failed to validate incorrect item")
		// return
	}

	mockItem = item{
		Amount: 100,
	}
	if _, ok := validateNewItem(mockItem); ok {
		t.Error("Failed to validate incorrect item")
		// return
	}
}

func TestCalculateTax(t *testing.T) {
	mockItem := item{
		TaxCode: 1,
		Amount:  100,
	}

	// Case 1
	var expected float64 = 10
	if result := calculateTax(mockItem); result != expected {
		t.Error("Wrong tax calculation - Case 1")
		return
	}

	// Case 2
	expected = 30
	mockItem.TaxCode = 2
	if result := calculateTax(mockItem); result != expected {
		t.Error("Wrong tax calculation - Case 2")
		return
	}

	// Case 3 - below/equal 100
	expected = 0
	mockItem.TaxCode = 3
	if result := calculateTax(mockItem); result != expected {
		t.Error("Wrong tax calculation - Case 3 - below/equal 100")
		return
	}

	// Case 3 - Above 100
	expected = 0.1
	mockItem.Amount = 101
	if result := calculateTax(mockItem); result != expected {
		t.Error("Wrong tax calculation - Case 3 - above 100")
		return
	}
}

func TestSetTaxType(t *testing.T) {
	expected := "Food"
	if result := setTaxType(1); result != expected {
		t.Error("Wrong set tax type")
	}

	expected = "Tobacco"
	if result := setTaxType(2); result != expected {
		t.Error("Wrong set tax type")
	}

	expected = "Entertainment"
	if result := setTaxType(3); result != expected {
		t.Error("Wrong set tax type")
	}

	expected = ""
	if result := setTaxType(4); result != expected {
		t.Error("Wrong set tax type")
	}
}

func TestCalculateTotal(t *testing.T) {
	mockItems := []item{}
	mockItem1 := item{
		TaxAmount: 10,
		Amount:    100,
	}
	mockItems = append(mockItems, mockItem1)
	mockItem2 := item{
		TaxAmount: 30,
		Amount:    300,
	}
	mockItems = append(mockItems, mockItem2)

	var expectedTaxAmount float64 = 40
	var expectedTotalAmount float64 = 400
	var expectedGrandTotal float64 = 440
	resultTaxAmount, resulTotalAmount, resultGrandTotal := calculateTotal(mockItems)
	if resultTaxAmount != expectedTaxAmount {
		t.Error("Wrong total calculation - tax amount")
	}
	if resulTotalAmount != expectedTotalAmount {
		t.Error("Wrong total calculation - total amount")
	}
	if resultGrandTotal != expectedGrandTotal {
		t.Error("Wrong total calculation - grand total")
	}
}

func TestInsertNewItem(t *testing.T) {
	mockDB, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")

	mockItem := item{
		Name:        "test_1",
		Amount:      100,
		TaxCode:     1,
		Type:        "Food",
		TaxAmount:   0.1,
		TotalAmount: 1.1,
	}

	if err := insertNewItem(mockDB, mockItem); err != nil {
		t.Error(err)
	}

	// Delete mock data from database
	if _, err := mockDB.Exec("DELETE FROM items where name='test_1'"); err != nil {
		t.Error(err)
	}
}

func TestRetrieveAllItems(t *testing.T) {
	mockDB, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")

	if _, err := retrieveAllItems(mockDB); err != nil {
		t.Error(err)
	}
}

func TestGetItems(t *testing.T) {
	mockDB, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")

	if _, err := getItems(mockDB); err != nil {
		t.Error(err)
	}
}
