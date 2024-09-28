package main

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/willf/bloom"
)

func main() {
	db, err := leveldb.OpenFile("mydb", nil)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer db.Close()

	bloomFilter := bloom.New(100000, 5)

	initialItems := []string{"key1", "key2", "key3"}

	for _, item := range initialItems {
		err = db.Put([]byte(item), []byte("value"), nil)
		if err != nil {
			log.Fatalf("Failed to put initial data: %v", err)
		}
		bloomFilter.Add([]byte(item))
		fmt.Printf("Added %s to the database and Bloom filter.\n", item)
	}

	newItems := []string{"key1", "key4"}

	for _, newItem := range newItems {
		if bloomFilter.Test([]byte(newItem)) {
			existingValue, err := db.Get([]byte(newItem), nil)
			_ = existingValue
			if err == nil {

				fmt.Printf("Updating %s in the database.\n", newItem)
				err = db.Put([]byte(newItem), []byte("existingValue"), nil)
				if err != nil {
					log.Fatalf("Failed to update data: %v", err)
				}
			}
		} else {

			fmt.Printf("Adding %s to the database.\n", newItem)
			err = db.Put([]byte(newItem), []byte("new_value"), nil)
			if err != nil {
				log.Fatalf("Failed to add new data: %v", err)
			}
			bloomFilter.Add([]byte(newItem))

		}
	}
}
