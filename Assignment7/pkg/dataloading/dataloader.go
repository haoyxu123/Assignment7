package dataloading

import (
	"encoding/csv"
	"os"
	"strconv"
)

type DataPoint struct {
	X, Y float64
}

type DataSet map[string][]DataPoint

func LoadCSV(filename string) (DataSet, error) {
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
		datasets[datasetName] = append(datasets[datasetName], DataPoint{X: x, Y: y})
	}
	return datasets, nil
}
