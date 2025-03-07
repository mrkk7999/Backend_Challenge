package implementation

import (
	"errors"
	"fmt"
	"strings"
)

func (s *service) Echo(matrix [][]int) (string, error) {
	if len(matrix) == 0 {
		return "", errors.New("matrix is empty")
	}

	cols := len(matrix[0])
	var result []string

	for _, row := range matrix {
		if len(row) != cols {
			return "", errors.New("matrix is not rectangular")
		}

		rowStr := make([]string, len(row))
		for i, val := range row {
			rowStr[i] = fmt.Sprintf("%d", val)
		}
		result = append(result, strings.Join(rowStr, ","))
	}

	return strings.Join(result, "\n"), nil
}

// Invert
// transposes the matrix
func (s *service) Invert(matrix [][]int) (string, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return "", errors.New("matrix is empty or contains empty rows")
	}

	rows, cols := len(matrix), len(matrix[0])

	// Ensure all rows have the same length
	for _, row := range matrix {
		if len(row) != cols {
			return "", errors.New("matrix is not rectangular")
		}
	}

	// Allocate a transposed matrix
	inverted := make([][]int, cols)
	for i := range inverted {
		inverted[i] = make([]int, rows)
	}

	// Perform transpose
	for i, row := range matrix {
		for j, val := range row {
			inverted[j][i] = val
		}
	}

	return s.Echo(inverted)
}

// Flatten
// converts the matrix to a single line
func (s *service) Flatten(matrix [][]int) (string, error) {
	if len(matrix) == 0 {
		return "", errors.New("matrix is empty")
	}

	var result []string
	for _, row := range matrix {
		if len(row) == 0 {
			return "", errors.New("matrix contains an empty row")
		}

		for _, val := range row {
			result = append(result, fmt.Sprintf("%d", val))
		}
	}
	return strings.Join(result, ","), nil
}

// Sum
// calculates the sum of all elements
func (s *service) Sum(matrix [][]int) (int, error) {
	if len(matrix) == 0 {
		return 0, errors.New("matrix is empty")
	}

	sum := 0
	for _, row := range matrix {
		if len(row) == 0 {
			return 0, errors.New("matrix contains an empty row")
		}

		for _, val := range row {
			sum += val
		}
	}
	return sum, nil
}

// Multiply
// calculates the product of all elements
func (s *service) Multiply(matrix [][]int) (int, error) {
	if len(matrix) == 0 {
		return 0, errors.New("matrix is empty")
	}

	product := 1
	for _, row := range matrix {
		if len(row) == 0 {
			return 0, errors.New("matrix contains an empty row")
		}

		for _, val := range row {
			product *= val
		}
	}
	return product, nil
}
