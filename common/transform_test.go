package common

import (
	"net/http/httptest"
	"testing"
)

func TestReadBody(t *testing.T) {
	mockRead := httptest.NewRequest("GET", "/item", nil)
	_, err := ReadBody(mockRead)
	if err != nil {
		t.Error(err)
	}
}

func TestNormalizeSpace(t *testing.T) {
	mockInput := "test      "
	expected := "test"
	result := NormalizeSpace(mockInput)
	if result != expected {
		t.Error("Failed to normalized space")
	}
}
