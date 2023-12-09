package helpers

import (
	"reflect"
	"testing"
)

func TestForEach(t *testing.T) {
	// Example slice of integers
	slice := []int{1, 2, 3}

	// Variable to track the function execution
	var sum int

	// Function to pass to ForEach
	fn := func(index int, value int) {
		sum += value
	}

	// Call ForEach
	ForEach(slice, fn)

	// Check if sum is as expected
	expectedSum := 6 // 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}
}

func TestMap(t *testing.T) {
	// Example slice of integers
	slice := []int{1, 2, 3}

	// Transformation function: multiply each element by 2
	fn := func(element int, index int, s []int) int {
		return element * 2
	}

	// Apply Map function
	result := Map(slice, fn)

	// Define the expected result
	expected := []int{2, 4, 6}

	// Compare the result with the expected result
	if len(result) != len(expected) {
		t.Fatalf("Expected slice length %d, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestFilter(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		slice    []int
		fn       func(int) bool
		expected []int
	}{
		{
			name:  "filter even numbers",
			slice: []int{1, 2, 3, 4, 5, 6},
			fn: func(n int) bool {
				return n%2 == 0
			},
			expected: []int{2, 4, 6},
		},
		{
			name:  "filter positive numbers",
			slice: []int{-2, -1, 0, 1, 2},
			fn: func(n int) bool {
				return n > 0
			},
			expected: []int{1, 2},
		},
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.slice, tt.fn); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Filter() = %v, want %v", got, tt.expected)
			}
		})
	}
}
