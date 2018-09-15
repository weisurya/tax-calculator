package config

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/middleware"
	cors "github.com/go-chi/cors"
	_ "github.com/lib/pq"
	yaml "gopkg.in/yaml.v2"

	common "tax-calculator/common"
	tax "tax-calculator/endpoint/tax"
)

// validateGoPATH function uses to validate the GOPATH and GOROOT of the environment whether it has been set up or not
func validateGoPATH() (err error) {
	// Check whether or not the GOPATH directory(ies) has been set up on the system
	if len(gopath) == 0 {
		err = errors.New("GOPATH environment value is empty. Please refer to http://golang.org/doc/code.html to configure your GO environment")
		return err
	}

	// Check whether or not the directory of GOPATH is the similar one with GOROOT
	if gopath == goroot {
		err = fmt.Errorf("GOPATH (%s) environment must be in different directory with GOROOT (%s). Please refer to http://golang.org/doc/code.html to configure your GO environment", gopath, goroot)
		return err
	}
	return nil
}

// getConfigurationFilePATH function uses to retrieve the directory path where the configuration file was stored
func getConfigurationFilePATH(projectPATH string) string {
	return filepath.Join(projectPATH, filepath.Join("storage", defautConfigurationFilename))
}

// loadConfiguration function uses to open up the configuration file based on the configPATH that was provided and unmarshal it based on config structure
func loadConfiguration(configPATH string) (config *Configuration, err error) {
	content, err := ioutil.ReadFile(configPATH)
	if err != nil {
		err = errors.New("Failed to read configuration file")
		return
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		err = errors.New("Failed to unmarshal YAML config information")
		return
	}

	return
}

// writeSystemURL function uses to write the project URL for provider
func writeSystemURL(host string, port int) string {
	var url bytes.Buffer
	url.WriteString(host)
	url.WriteString(":")
	url.WriteString(strconv.Itoa(port))
	return url.String()
}

// initializeDatabase function uses to initialize a database connection
func initializeDatabase(config database) (db *sql.DB, err error) {
	databaseURL := writeDatabaseURL(config)
	db, err = sql.Open(config.Connection, databaseURL)
	if err != nil {
		err = errors.New(" Failed to connect the " + config.Connection + " database, " + err.Error())
		return
	}
	if err = db.Ping(); err != nil {
		err = errors.New(" Failed to ping the " + config.Connection + " database, " + err.Error())
		return
	}
	db.SetMaxOpenConns(defaultMaxOpenConnection)
	fmt.Println("Successfully connected to database: " + config.Name)
	return
}

// writeDatabaseURL function uses to write the command to connect database endpoint
func writeDatabaseURL(config database) string {
	var database bytes.Buffer
	portString := strconv.Itoa(config.Port)
	database.WriteString(config.Connection + "://")
	database.WriteString(config.User + ":" + config.Password)
	database.WriteString("@")
	database.WriteString(config.Host + ":" + portString + "/")
	database.WriteString(config.Name)
	database.WriteString("?sslmode=disable")
	return database.String()
}

// initializeRouterConfiguration function uses to initialize router configuration
func initializeRouterConfiguration() *chi.Mux {
	router := chi.NewRouter()
	router.Use(setCORS().Handler)
	router.Use(middleware.Logger)
	return router
}

// setCORS function uses to set up the CORS for the router configuration
func setCORS() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Requested-With", "access-token", "initial-code"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	return cors
}

// setURIList function uses as the function to define the endpoint API of this system
func setURIList(r *chi.Mux) {
	r.MethodNotAllowed(common.MethodNotAllowedHandler())

	r.Route("/item", func(r chi.Router) {
		r.Get("/", tax.GetTotal)
		r.Post("/", tax.PostItem)
	})
}
