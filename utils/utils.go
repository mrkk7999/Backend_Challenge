package utils

import (
	"errors"
	"strconv"
)

// ParseCSV 
// converts CSV records to a 2D integer matrix
func ParseCSV(records [][]string) ([][]int, error) {
	if len(records) == 0 {
		return nil, errors.New("empty file")
	}

	size := len(records)
	matrix := make([][]int, size)

	for i, row := range records {
		if len(row) != size {
			return nil, errors.New("matrix must be square")
		}
		matrix[i] = make([]int, size)
		for j, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("non-integer value found")
			}
			matrix[i][j] = num
		}
	}
	return matrix, nil
}
