package common

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateInsertStatement(t *testing.T) {
	mockTable := "item"
	mockField := "name, amount"
	mockValue := "'test', 1"
	expected := "INSERT INTO items(name, amount) VALUES('test', 1);"
	result := CreateInsertStatement(mockTable, mockField, mockValue)
	if result != expected {
		t.Error("Failed on creating insert statement")
	}
}

func TestExecuteSQL(t *testing.T) {
	mockDB, _ := sql.Open("postgres", "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable")

	mockInsert := "INSERT INTO items(name, amount, taxcode, type, taxamount, totalamount) VALUES('test', 100, 1, 'Food', 1.1, 1.1);"
	if err := ExecuteSQL(mockDB, mockInsert); err != nil {
		t.Error(err)
	}

	mockDelete := "DELETE FROM items WHERE name='test'"
	if err := ExecuteSQL(mockDB, mockDelete); err != nil {
		t.Error(err)
	}
}
