package main

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	// Initialize Bloom filter with a capacity of 100,000 and 5 hash functions
	filter := bloom.New(100000, 5)

	// Simulate receiving data in a continuous stream
	stream := []string{"user1", "user2", "user3", "user4", "user1"}

	// Test and add each item to the filter, handling duplicates efficiently
	for _, item := range stream {
		if filter.TestAndAdd([]byte(item)) {
			fmt.Printf("%s was already present. Skipping...\n", item)
		} else {
			fmt.Printf("%s added to the filter.\n", item)
		}
	}
}
