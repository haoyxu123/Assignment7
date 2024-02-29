package main

import (
	"fmt"
	"log"
	"time"

	"example.com/assignment3/pkg/dataloading"
	"example.com/assignment3/pkg/regression"
)

func main() {
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		datasets, err := dataloading.LoadCSV("C:/Assignment2/dataset/anscombes.csv")
		if err != nil {
			log.Fatalf("Error loading CSV file: %v", err)
		}
		for name, data := range datasets {
			if err := regression.PlotRegression(name, data); err != nil {
				log.Fatalf("could not plot dataset %s: %v", name, err)
			}
		}
	}
	elapsed := time.Now().Sub(startTime)
	fmt.Printf("Execution took %v\n", elapsed)
}
