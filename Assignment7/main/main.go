package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type DataPoint struct {
	x, y float64
}

type DataSet map[string][]DataPoint

func main() {
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		datasets, err := loadCSV("C:/Assignment2/dataset/anscombes.csv")
		if err != nil {
			log.Fatalf("Error loading CSV file: %v", err)
		}

		for name, data := range datasets {
			if err := plotRegression(name, data); err != nil {
				log.Fatalf("could not plot dataset %s: %v", name, err)
			}
		}
	}
	elapsed := time.Now().Sub(startTime)
	fmt.Printf("Execution took %d nanoseconds\n", elapsed)
}

func loadCSV(filename string) (DataSet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	datasets := make(DataSet)
	for _, record := range records[1:] {
		datasetName := record[1]
		x, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}
		datasets[datasetName] = append(datasets[datasetName], DataPoint{x: x, y: y})
	}
	return datasets, nil
}

func plotRegression(name string, data []DataPoint) error {
	p := plot.New()

	p.Title.Text = "Linear Regression for " + name
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	pts := make(plotter.XYs, len(data))
	for i := range data {
		pts[i].X = data[i].x
		pts[i].Y = data[i].y
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	// Assuming pts is of type plotter.XYs
	xs := make([]float64, len(pts))
	ys := make([]float64, len(pts))
	for i, pt := range pts {
		xs[i] = pt.X
		ys[i] = pt.Y
	}

	alpha, beta := stat.LinearRegression(xs, ys, nil, false)

	line := plotter.NewFunction(func(x float64) float64 {
		return beta*x + alpha
	})
	line.Color = color.RGBA{B: 255, A: 255}
	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, name+"_regression.png"); err != nil {
		return err
	}

	return nil
}
