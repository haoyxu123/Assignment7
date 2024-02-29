package main

import (
	"log"

	"example.com/assignment3/pkg/dataloading"
	"example.com/assignment3/pkg/regression"
)

func main() {
	filePath := "C:/Assignment2/dataset/anscombes.csv"

	datasets, err := dataloading.LoadCSV(filePath)
	if err != nil {
		log.Fatalf("Error loading CSV file: %v", err)
	}

	for name, data := range datasets {
		if err := regression.PlotRegression(name, data); err != nil {
			log.Fatalf("could not plot dataset %s: %v", name, err)
		}
	}
}
