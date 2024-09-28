package main

import (
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

}
