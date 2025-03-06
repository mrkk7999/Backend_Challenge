package iface

type Service interface {
	Echo(matrix [][]int) string
	Invert(matrix [][]int) string
	Flatten(matrix [][]int) string
	Sum(matrix [][]int) int
	Multiply(matrix [][]int) int
}
