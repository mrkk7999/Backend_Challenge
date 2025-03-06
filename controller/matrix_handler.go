package controller

import (
	"backend_challenge/utils"
	"encoding/csv"
	"fmt"
	"net/http"
)

// Generic function to handle requests
func (c *Controller) handleRequest(w http.ResponseWriter, r *http.Request, operation func([][]int) (interface{}, error)) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("File error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("CSV parsing error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix, err := utils.ParseCSV(records)
	if err != nil {
		http.Error(w, fmt.Sprintf("Matrix error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := operation(matrix)
	if err != nil {
		http.Error(w, fmt.Sprintf("Processing error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, result)
}

func (c *Controller) Echo(w http.ResponseWriter, r *http.Request) {
	c.handleRequest(w, r, func(matrix [][]int) (interface{}, error) {
		return c.service.Echo(matrix)
	})
}

func (c *Controller) Invert(w http.ResponseWriter, r *http.Request) {
	c.handleRequest(w, r, func(matrix [][]int) (interface{}, error) {
		return c.service.Invert(matrix)
	})
}

func (c *Controller) Flatten(w http.ResponseWriter, r *http.Request) {
	c.handleRequest(w, r, func(matrix [][]int) (interface{}, error) {
		return c.service.Flatten(matrix)
	})
}

func (c *Controller) Sum(w http.ResponseWriter, r *http.Request) {
	c.handleRequest(w, r, func(matrix [][]int) (interface{}, error) {
		return c.service.Sum(matrix)
	})
}

func (c *Controller) Multiply(w http.ResponseWriter, r *http.Request) {
	c.handleRequest(w, r, func(matrix [][]int) (interface{}, error) {
		return c.service.Multiply(matrix)
	})
}
