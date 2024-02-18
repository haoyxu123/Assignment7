package main

import (
	"fmt"
	"testing"
)

func TestLoadCSV(t *testing.T) {
	datasets, err := loadCSV("C:/Assignment2/dataset/anscombes.csv")
	if err != nil {
		t.Errorf("loadCSV() returned an error: %v", err)
	} else {
		fmt.Printf("datasets: %+v\n", datasets)
	}
}
