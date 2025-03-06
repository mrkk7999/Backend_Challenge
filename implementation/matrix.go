package implementation

import (
	"fmt"
	"strings"
)

// Echo
// returns the original matrix as a string
func (s *service) Echo(matrix [][]int) string {
	var result []string
	for _, row := range matrix {
		rowStr := make([]string, len(row))
		for i, val := range row {
			rowStr[i] = fmt.Sprintf("%d", val)
		}
		result = append(result, strings.Join(rowStr, ","))
	}
	return strings.Join(result, "\n")
}

// Invert
// transposes the matrix
func (s *service) Invert(matrix [][]int) string {

	size := len(matrix)
	inverted := make([][]int, size)
	for i := range inverted {
		inverted[i] = make([]int, size)
	}
	for i := range inverted {
		for j := range matrix[i] {
			inverted[j][i] = matrix[i][j]
		}
	}

	return s.Echo(inverted)
}

// Flatten
// converts the matrix to a single line
func (s *service) Flatten(matrix [][]int) string {

	var result []string
	for _, row := range matrix {
		for _, val := range row {
			result = append(result, fmt.Sprintf("%d", val))
		}
	}
	return strings.Join(result, ",")
}

// Sum
// calculates the sum of all elements
func (s *service) Sum(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

// Multiply
// calculates the product of all elements
func (s *service) Multiply(matrix [][]int) int {
	product := 1
	for _, row := range matrix {
		for _, val := range row {
			product *= val
		}
	}
	return product
}
