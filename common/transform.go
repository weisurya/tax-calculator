// Package common - transform. This file contains several functions that use to transform a oarticular variable into different output
package common

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// ReadBody functions uses to read the body request and transform it into byte
func ReadBody(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return body, err
	}
	defer r.Body.Close()
	return body, nil
}

// NormalizeSpace function uses to trim exceed space at particular variable
func NormalizeSpace(input string) (output string) {
	output = strings.Replace(input, " ", "", -1)
	return
}
