package main

import (
	"fmt"
	"time"

	"github.com/willf/bloom"
)

func main() {
	start := time.Now()
	var countFound int
	var countNotFound int
	memStore := make(map[string]int)
	// Initialize a Bloom filter with a capacity of 1 million and 10 hash functions
	filter := bloom.New(41_000_000, 13)
	// Simulate a dataset containing one million records
	for i := 0; i < 2_000_000; i++ {
		key := fmt.Sprintf("item%d", i)
		filter.Add([]byte(key))
		memStore[key] = i
	}

	/*n := filter.Test([]byte(fmt.Sprintf("item%d", 1_000_001)))
	_ = n
	return*/

	// Check if a few items are present in the filter

	for i := 0; i < 20_000_000; i++ {
		test := fmt.Sprintf("item%d", i)
		if filter.Test([]byte(test)) {
			if _, ok := memStore[test]; !ok {
				fmt.Println("Exists But Failed for ", test)
				countFound++
			}
		} else {
			if _, ok := memStore[test]; ok {
				fmt.Println("No Exists But Failed for ", test)
			}
			countNotFound++
		}
	}

	fmt.Println("Done!!")
	fmt.Println(countFound)
	fmt.Println(countNotFound)
	duration := time.Since(start)
	fmt.Println(duration)
}
