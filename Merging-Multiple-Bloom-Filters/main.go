package main

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	// Create two Bloom filters
	bf1 := bloom.New(100000, 5)
	bf2 := bloom.New(100000, 5)

	// Add data to the first filter
	items1 := []string{"apple", "orange", "banana"}
	for _, item := range items1 {
		bf1.Add([]byte(item))
	}

	// Add data to the second filter
	items2 := []string{"grape", "apple", "pear"}
	for _, item := range items2 {
		bf2.Add([]byte(item))
	}

	// Merge the second filter into the first one
	bf1.Merge(bf2)

	// Test data after merging
	tests := []string{"apple", "banana", "grape", "watermelon"}
	for _, test := range tests {
		if bf1.Test([]byte(test)) {
			fmt.Printf("%s might be in the filter.\n", test)
		} else {
			fmt.Printf("%s is definitely not in the filter.\n", test)
		}
	}
}
