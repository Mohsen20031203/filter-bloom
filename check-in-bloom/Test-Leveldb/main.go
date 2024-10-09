package main

import (
	"fmt"
	"time"

	"github.com/willf/bloom"
)

func main() {
	var countFound float64
	var countNotFound float64
	memStore := make(map[string]int)
	// Initialize a Bloom filter with a capacity of 1 million and 10 hash functions
	filter := bloom.New(65_000_000, 10)
	// Simulate a dataset containing one million records
	for i := 0; i <= 2_000_000; i++ {
		key := fmt.Sprintf("item%d", i)
		filter.Add([]byte(key))
		memStore[key] = i
	}

	// Check if a few items are present in the filter
	start := time.Now()

	var numberAll float64
	for i := 2_000_000; i < 3_000_000; i++ {
		numberAll++
		test := fmt.Sprintf("item%d", i)
		if filter.Test([]byte(test)) {
			if _, ok := memStore[test]; !ok {
				countFound++
			}
		} else {
			if _, ok := memStore[test]; ok {
				fmt.Println("No Exists But Failed for ", test)
				countNotFound++
			}
		}
	}

	fmt.Println("bloom = t | map= f : ", countFound)
	fmt.Println("bloom = f | map= f : ", countNotFound)
	duration := time.Since(start)
	fmt.Println("All item for check : ", numberAll)
	fmt.Println("Error percentage‌   : ", (countFound*100)/numberAll)
	fmt.Println("time check in data : ", duration)
	fmt.Println("Done!!")
}
