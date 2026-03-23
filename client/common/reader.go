package common

import (
	"encoding/csv"
	"os"
)

// Reads the CSV file from the path passed as parameter
func read_csv_file(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
