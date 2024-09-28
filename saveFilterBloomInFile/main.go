package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/willf/bloom"
)

func main() {
	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Create a Bloom filter with capacity for 1,000,000 items and 5 hash functions
	bf := bloom.New(1000000, 5)

	// Add random data to the filter
	for i := 0; i < 1000; i++ {
		randomValue := fmt.Sprintf("item-%d", rand.Intn(1000000))
		bf.Add([]byte(randomValue))
		fmt.Printf("Added %s to Bloom filter\n", randomValue)
	}

	// Save the Bloom filter to a file
	file, err := os.Create("bloomfilter.dat")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = bf.WriteTo(file)
	if err != nil {
		fmt.Printf("Error writing Bloom filter to file: %v\n", err)
	}
}
