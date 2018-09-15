// Package provider contains several function that used to provide several endpoint/variable that intentionally used at API layer. This package becomes the middleware between package config and API layer. It was separated to avoid error import cycle.
package provider

import (
	"database/sql"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

// provider struct defines list of variable that was provider for API layer
type provider struct {
	URL      string
	Database *sql.DB
	Router   *chi.Mux
}

// set global variable that use as the provider variable
var Provider provider

// UpdateURL function uses to update the latest URL on provider
func UpdateURL(url string) {
	Provider.URL = url
}

// UpdateDatabase function uses to update the latest database endpoint on provider
func UpdateDatabase(db *sql.DB) {
	Provider.Database = db
}

// UpdateRouter function uses to update the latest router endpoint on provider
func UpdateRouter(router *chi.Mux) {
	Provider.Router = router
}
