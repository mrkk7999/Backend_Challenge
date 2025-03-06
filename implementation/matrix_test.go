package implementation

import (
	"testing"
)

func TestEcho(t *testing.T) {
	service := &service{}
	tests := []struct {
		name      string
		matrix    [][]int
		expected  string
		expectErr bool
	}{
		{"Echo Empty Matrix", [][]int{}, "", true},
		{"Echo Non-Rectangular Matrix", [][]int{{1, 2}, {3}}, "", true},
		{"Echo Valid Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,2,3\n4,5,6\n7,8,9", false},
		{"Echo Single Element Matrix", [][]int{{5}}, "5", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Echo(tt.matrix)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	service := &service{}
	tests := []struct {
		name      string
		matrix    [][]int
		expected  string
		expectErr bool
	}{
		{"Invert Empty Matrix", [][]int{}, "", true},
		{"Invert Non-Rectangular Matrix", [][]int{{1, 2}, {3}}, "", true},
		{"Invert Valid Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,4,7\n2,5,8\n3,6,9", false},
		{"Invert Single Element Matrix", [][]int{{5}}, "5", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Invert(tt.matrix)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	service := &service{}
	tests := []struct {
		name      string
		matrix    [][]int
		expected  string
		expectErr bool
	}{
		{"Flatten Empty Matrix", [][]int{}, "", true},
		{"Flatten Valid Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "1,2,3,4,5,6,7,8,9", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Flatten(tt.matrix)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestSum(t *testing.T) {
	service := &service{}
	tests := []struct {
		name      string
		matrix    [][]int
		expected  int
		expectErr bool
	}{
		{"Sum Empty Matrix", [][]int{}, 0, true},
		{"Sum Valid Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 45, false},
		{"Sum Zero Matrix", [][]int{{0, 0}, {0, 0}}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Sum(tt.matrix)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	service := &service{}
	tests := []struct {
		name      string
		matrix    [][]int
		expected  int
		expectErr bool
	}{
		{"Multiply Empty Matrix", [][]int{}, 0, true},
		{"Multiply Valid Matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 362880, false},
		{"Multiply Zero Matrix", [][]int{{0, 0}, {0, 0}}, 0, false},
		{"Multiply Single Element Matrix", [][]int{{5}}, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Multiply(tt.matrix)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
