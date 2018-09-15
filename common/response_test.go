package common

import (
	"net/http/httptest"
	"testing"
	"time"
)

func TestMethodNotAllowedHandler(t *testing.T) {
	MethodNotAllowedHandler()
}

func TestCreateResponse(t *testing.T) {
	http := &HTTP{
		Write: httptest.NewRecorder(),
		Read:  httptest.NewRequest("GET", "/item", nil),
	}

	mock2xxCode := NoErrorCode
	mock400Code := ErrorBadRequestCode
	mock405Code := ErrorMethodNotAllowedCode
	mock404Code := ErrorNotFound
	mock422Code := ErrorUnprocessableEntityCode
	mock500Code := ErrorInternalServerCode

	var emptyData interface{}
	containData := "mock data"

	CreateResponse(*http, mock2xxCode, containData)
	CreateResponse(*http, mock400Code, containData)
	CreateResponse(*http, mock405Code, containData)
	CreateResponse(*http, mock404Code, emptyData)
	CreateResponse(*http, mock422Code, emptyData)
	CreateResponse(*http, mock500Code, emptyData)
}

func TestWriteResponseHeader(t *testing.T) {
	http := &HTTP{
		Write: httptest.NewRecorder(),
		Read:  httptest.NewRequest("GET", "/item", nil),
	}

	mock2xxCode := NoErrorCode
	mock400Code := ErrorBadRequestCode
	mock405Code := ErrorMethodNotAllowedCode
	mock404Code := ErrorNotFound
	mock422Code := ErrorUnprocessableEntityCode
	mock500Code := ErrorInternalServerCode

	writeResponseHeader(*http, mock2xxCode)
	writeResponseHeader(*http, mock400Code)
	writeResponseHeader(*http, mock405Code)
	writeResponseHeader(*http, mock404Code)
	writeResponseHeader(*http, mock422Code)
	writeResponseHeader(*http, mock500Code)
}

func TestWriteResponseBody(t *testing.T) {
	http := &HTTP{
		Write: httptest.NewRecorder(),
		Read:  httptest.NewRequest("GET", "/item", nil),
	}

	mock2xxCode := NoErrorCode
	mock400Code := ErrorBadRequestCode
	mock405Code := ErrorMethodNotAllowedCode
	mock404Code := ErrorNotFound
	mock422Code := ErrorUnprocessableEntityCode
	mock500Code := ErrorInternalServerCode
	var emptyData interface{}
	containData := "mock data"
	mockTime := time.Now().Format(time.RFC3339)

	writeResponseBody(*http, mock2xxCode, emptyData, mockTime)
	writeResponseBody(*http, mock400Code, emptyData, mockTime)
	writeResponseBody(*http, mock405Code, emptyData, mockTime)
	writeResponseBody(*http, mock404Code, emptyData, mockTime)
	writeResponseBody(*http, mock422Code, emptyData, mockTime)
	writeResponseBody(*http, mock500Code, emptyData, mockTime)

	writeResponseBody(*http, mock2xxCode, containData, mockTime)
	writeResponseBody(*http, mock400Code, containData, mockTime)
	writeResponseBody(*http, mock405Code, containData, mockTime)
	writeResponseBody(*http, mock404Code, containData, mockTime)
	writeResponseBody(*http, mock422Code, containData, mockTime)
	writeResponseBody(*http, mock500Code, containData, mockTime)
}

func TestWriteTimeIn(t *testing.T) {
	if result := writeTimeIn(); result == "" {
		t.Error("Failed to write time in")
	}
}
