package common

import (
	"encoding/json"
	"net/http"
	"time"
)

// List of standarize error code
type errorCode int

const (
	// NoErrorCode - 200
	NoErrorCode errorCode = http.StatusOK
	// ErrorBadRequestCode - 400
	ErrorBadRequestCode errorCode = http.StatusBadRequest
	// ErrorMethodNotAllowedCode - 500
	ErrorMethodNotAllowedCode errorCode = http.StatusMethodNotAllowed
	// ErrorNotFound - 404
	ErrorNotFound errorCode = http.StatusNotFound
	// ErrorUnprocessableEntityCode - 422
	ErrorUnprocessableEntityCode errorCode = http.StatusUnprocessableEntity
	// ErrorInternalServerCode - 500
	ErrorInternalServerCode errorCode = http.StatusInternalServerError
)

const (
	// defaultTimeFormat - "02-01-2006 15:04:05"
	defaultTimeFormat = time.RFC3339
)

// HTTP struct defines the write and read struct from net/http
type HTTP struct {
	Write http.ResponseWriter
	Read  *http.Request
}

// writeResponseHeader function uses to create a standarize header response
func writeResponseHeader(HTTP HTTP, code errorCode) {
	HTTP.Write.Header().Set("Content-Type", "application/json")
	HTTP.Write.WriteHeader(int(code))
}

// writeResponseBody function uses to create a standarize body response
func writeResponseBody(HTTP HTTP, code errorCode, data interface{}, timeIn string) {
	if data == nil {
		switch code {
		case NoErrorCode:
			data = "OK"
		case ErrorBadRequestCode:
			data = "Bad request"
		case ErrorMethodNotAllowedCode:
			data = "Method not allowed"
		case ErrorNotFound:
			data = "Not found"
		case ErrorUnprocessableEntityCode:
			data = "Unable to process the request"
		case ErrorInternalServerCode:
			data = "Internal server error"
		}
	}

	body, _ := json.Marshal(data)
	HTTP.Write.Write(body)
}

// writeTimeIn function uses to create a standarize time in format
func writeTimeIn() string {
	return time.Now().Format(defaultTimeFormat)
}
