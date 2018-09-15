// Package main stands as the main package to utilize the system
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	config "tax-calculator/config"
)

// List of exit code
const (
	cleanExit = iota
	errorExit
	forceExit
)

// A Group of flag variables that is used to list the available flag for this system
var (
	needHelp = flag.Bool("help", false, "Print help information")
	host     = flag.String("host", "", "Host name to access system ('host:port')")
	port     = flag.Int("port", 9090, "Port number to access system ('host:port')")
)

func init() {
	// Set the maximum number of CPUs for this system as equal to the number of CPUs, that available in the environment where this system will be initialized, as default
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// main function always listen the return value of exit code
func main() {
	go os.Exit(initializeSystem())
}

// initializeSystem function defines the main function to start up the system through command-line
func initializeSystem() int {
	if checkFlag(); *needHelp {
		showHelp()
		return cleanExit
	}

	projectRootPATH, err := getProjectRootPATH()
	if err != nil {
		log.Println(err)
		return errorExit
	}

	configuration, err := config.InitializeConfiguration(projectRootPATH)
	if err != nil {
		log.Println(err)
		return errorExit
	}

	if err := config.UtilizeSystem(configuration); err != nil {
		log.Println(err)
		return errorExit
	}

	return cleanExit
}

// checkFlag function uses to check the argument that was parsed as flag through command-line
func checkFlag() {
	flag.Parse()
}

// showHelp function uses to show the default settings of all defined command-line flags
func showHelp() {
	fmt.Println("List of available flags: ")
	flag.PrintDefaults()
}

// getProjectRootPATH function uses to retrieve the project directory as the root PATH for this system
func getProjectRootPATH() (projectPATH string, err error) {

	projectPATH, err = os.Getwd()
	if err != nil {
		return "", err
	}
	return projectPATH, nil
}
