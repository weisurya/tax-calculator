package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
)

// getProjectDirectory is a helper function to create the project PATH for testing purpose
func getProjectDirectory() string {
	currentPATH, _ := os.Getwd()
	// Modify the output of project PATH element by deleting the last element of path for testing purpose
	projectPATH := filepath.Dir(currentPATH)
	return projectPATH
}
func TestInitializeConfiguration(t *testing.T) {
	expected := &Configuration{
		Application: application{
			Name: "Tax Calculator",
			Host: "localhost",
			Port: 9090,
		},
		Database: database{
			Name:       "tax_calculator",
			User:       "postgres",
			Password:   "postgres",
			Port:       5432,
			Connection: "postgres",
			Host:       "tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com",
		},
	}

	projectPATH := getProjectDirectory()

	result, err := InitializeConfiguration(projectPATH)
	if err != nil {
		t.Error(err)
		return
	}

	if passed := reflect.DeepEqual(result, expected); !passed {
		t.Error("Failed to initialize configuration")
		return
	}
}

func TestValidateGoPATH(t *testing.T) {
	if err := validateGoPATH(); err != nil {
		t.Error(err)
	}
}

func TestGetConfigurationFilePATH(t *testing.T) {
	projectPATH := getProjectDirectory()

	if result := getConfigurationFilePATH(projectPATH); result == "" {
		t.Error("Failed to get configuration file PATH")
	}
}

func TestLoadConfiguration(t *testing.T) {
	projectPATH := getProjectDirectory()

	configPATH := getConfigurationFilePATH(projectPATH)
	if configPATH == "" {
		t.Error("Failed to get configuration file PATH")
	}

	if _, err := loadConfiguration(configPATH); err != nil {
		t.Error("Failed to load configuration")
	}
}

func TestWriteSystemURL(t *testing.T) {
	host := "localhost"
	port := 9090
	expected := host + ":" + strconv.Itoa(port)
	result := writeSystemURL(host, port)
	if result != expected {
		t.Error("Failed to write system URL")
	}
}

func TestInitializeDatabase(t *testing.T) {
	config := &Configuration{
		Database: database{
			Name:       "tax_calculator",
			User:       "postgres",
			Password:   "postgres",
			Port:       5432,
			Connection: "postgres",
			Host:       "tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com",
		},
	}

	if _, err := initializeDatabase(config.Database); err != nil {
		t.Error(err)
	}
}

func TestWriteDatabaseURL(t *testing.T) {
	config := &Configuration{
		Database: database{
			Name:       "tax_calculator",
			User:       "postgres",
			Password:   "postgres",
			Port:       5432,
			Connection: "postgres",
			Host:       "tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com",
		},
	}

	expected := "postgres://postgres:postgres@tax-calculator.cqveqnulaatq.us-west-2.rds.amazonaws.com:5432/tax_calculator?sslmode=disable"
	if result := writeDatabaseURL(config.Database); result != expected {
		t.Error("Failed to write database URL")
	}
}

func TestInitializeRouterConfiguration(t *testing.T) {
	if router := initializeRouterConfiguration(); router == nil {
		t.Error("Failed to initialize router configuration")
	}
}

func TestSetCORS(t *testing.T) {
	if cors := setCORS(); cors == nil {
		t.Error("Failed to set CORS")
	}
}

func TestSetURIList(t *testing.T) {
	router := initializeRouterConfiguration()
	setURIList(router)
}
