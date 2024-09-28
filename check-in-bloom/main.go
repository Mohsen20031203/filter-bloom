package main

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	// Initialize a Bloom filter with a capacity of 1 million and 10 hash functions
	filter := bloom.New(1_000_000, 10)

	// Simulate a dataset containing one million records
	for i := 0; i < 1_000_000; i++ {
		filter.Add([]byte(fmt.Sprintf("item%d", i)))
	}
}
