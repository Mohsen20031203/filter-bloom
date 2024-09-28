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

	// Check if a few items are present in the filter
	tests := []string{"item500", "item150000", "non_existent_item"}
	for _, test := range tests {
		if filter.Test([]byte(test)) {
			fmt.Printf("%s is possibly in the set.\n", test)
		} else {
			fmt.Printf("%s is definitely not in the set.\n", test)
		}
	}
}
