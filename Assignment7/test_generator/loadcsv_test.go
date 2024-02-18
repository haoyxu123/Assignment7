package main

import (
	"fmt"
	"testing"

	assignment3main "example.com/assignment3-main"
)

func TestLoadCSV(t *testing.T) {
	datasets, err := assignment3main.LoadCSV("C:/Assignment2/dataset/anscombes.csv")
	if err != nil {
		t.Errorf("loadCSV() returned an error: %v", err)
		fmt.Println("datasets")
	}
}
