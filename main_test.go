package main

import (
	"testing"
)

func TestCheckFlag(t *testing.T) {
	checkFlag()
}

func TestShowHelp(t *testing.T) {
	showHelp()
}

func TestGetProjectRootPATH(t *testing.T) {
	_, err := getProjectRootPATH()
	if err != nil {
		t.Error(err)
	}
}
