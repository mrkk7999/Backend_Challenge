package iface

type Service interface {
	Echo(matrix [][]int) (string, error)
	Invert(matrix [][]int) (string, error)
	Flatten(matrix [][]int) (string, error)
	Sum(matrix [][]int) (int, error)
	Multiply(matrix [][]int) (int, error)
}
