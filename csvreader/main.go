// Package csvreader reads CSV files
package csvreader

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Line represent each line in the wordlist
type Line struct {
	ID   int
	Word string
}

// Read reads a csv file and returns its content
func Read(filename string) ([]Line, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return []Line{}, err
	}
	defer f.Close()

	// Read file
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []Line{}, err
	}

	// Convert each line into a Line struct
	var result []Line
	for _, line := range lines {
		id, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			panic("Error reading file")
		}

		newLine := Line{
			ID:   int(id),
			Word: line[1],
		}

		result = append(result, newLine)
	}

	// Return file content as []Line
	return result, nil
}
