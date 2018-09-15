// Package common - response. This file contains several function that was intended to be used to create a standarize return response at API layer.
package common

import (
	"errors"
	"net/http"
)

// MethodNotAllowedHandler uses to handle whether the requester is not allowed to access API
func MethodNotAllowedHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errors.New("Method Not Allowed")
		http := HTTP{
			Write: w,
			Read:  r,
		}
		CreateResponse(http, ErrorMethodNotAllowedCode, err.Error())
	})
}

// CreateResponse function uses to create a standarize response at API layer
func CreateResponse(HTTP HTTP, code errorCode, data interface{}) {
	writeResponseHeader(HTTP, code)
	writeResponseBody(HTTP, code, data, writeTimeIn())
}
