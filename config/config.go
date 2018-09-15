// Package config contains several functions that was used for system configuration purpose
package config

import (
	"fmt"
	"go/build"
	"net/http"

	provider "tax-calculator/provider"
)

const (
	// the default configuration filename for this system
	defautConfigurationFilename = "credential.yml"

	// the default maximum open connection for postgreSQL
	defaultMaxOpenConnection = 20
)

// List of common variables that will be used in several parts of this package. Furthermore, it was initialized for readability purpose
var (
	gopath = build.Default.GOPATH
	goroot = build.Default.GOROOT
)

// Configuration structure defines the output result of the configuration
type Configuration struct {
	Application application `yaml:"application,omitempty"`
	Database    database    `yaml:"database,omitempty"`
}

// application structure helps to define the key value of application
type application struct {
	Name string `yaml:"name,omitempty"`
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

// database structure helps to define the key value of database
type database struct {
	Name       string `yaml:"name,omitempty"`
	User       string `yaml:"user,omitempty"`
	Password   string `yaml:"password,omitempty"`
	Port       int    `yaml:"port,omitempty"`
	Connection string `yaml:"connection,omitempty"`
	Host       string `yaml:"host,omitempty"`
}

// InitializeConfiguration function uses to initialize the required configuration to utilize this system
func InitializeConfiguration(projectRootPATH string) (configuration *Configuration, err error) {
	if err = validateGoPATH(); err != nil {
		return nil, err
	}

	configurationFilePATH := getConfigurationFilePATH(projectRootPATH)

	configuration, err = loadConfiguration(configurationFilePATH)
	if err != nil {
		return nil, err
	}

	url := writeSystemURL(configuration.Application.Host, configuration.Application.Port)
	provider.UpdateURL(url)

	db, err := initializeDatabase(configuration.Database)
	if err != nil {
		return nil, err
	}
	provider.UpdateDatabase(db)

	router := initializeRouterConfiguration()
	provider.UpdateRouter(router)

	return configuration, nil
}

// UtilizeSystem function uses to utilize the configuration that has been successfully initialized
func UtilizeSystem(config *Configuration) (err error) {
	fmt.Println("Listening to " + provider.Provider.URL)
	fmt.Println("Port type: HTTP")
	setURIList(provider.Provider.Router)

	if err := http.ListenAndServe(provider.Provider.URL, provider.Provider.Router); err != nil {
		return err
	}
	return nil
}
