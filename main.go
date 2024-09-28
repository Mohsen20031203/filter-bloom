package main

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/willf/bloom"
)

func main() {
	// باز کردن یا ایجاد دیتابیس
	db, err := leveldb.OpenFile("mydb", nil)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer db.Close()

	// ایجاد فیلتر بلوم با ظرفیت 100000 و 5 تابع هش
	bloomFilter := bloom.New(100000, 5)

	// داده‌های اولیه که می‌خواهیم به دیتابیس اضافه کنیم
	initialItems := []string{"key1", "key2", "key3"}

	// افزودن داده‌ها به دیتابیس و فیلتر بلوم
	for _, item := range initialItems {
		err = db.Put([]byte(item), []byte("value"), nil)
		if err != nil {
			log.Fatalf("Failed to put initial data: %v", err)
		}
		bloomFilter.Add([]byte(item)) // افزودن کلید به فیلتر بلوم
		fmt.Printf("Added %s to the database and Bloom filter.\n", item)
	}

	// مقادیر جدیدی که می‌خواهیم بررسی کنیم
	newItems := []string{"key1", "key4"}

	// بررسی و افزودن یا به‌روزرسانی مقادیر جدید
	for _, newItem := range newItems {
		if bloomFilter.Test([]byte(newItem)) { // چک کردن با فیلتر بلوم
			existingValue, err := db.Get([]byte(newItem), nil)
			_ = existingValue
			if err == nil {
				// اگر کلید موجود بود، مقدار آن را به‌روزرسانی می‌کنیم
				fmt.Printf("Updating %s in the database.\n", newItem)
				err = db.Put([]byte(newItem), []byte("existingValue"), nil)
				if err != nil {
					log.Fatalf("Failed to update data: %v", err)
				}
			}
		} else {
			// اگر کلید موجود نبود، آن را به دیتابیس اضافه می‌کنیم
			fmt.Printf("Adding %s to the database.\n", newItem)
			err = db.Put([]byte(newItem), []byte("new_value"), nil)
			if err != nil {
				log.Fatalf("Failed to add new data: %v", err)
			}
			bloomFilter.Add([]byte(newItem)) // افزودن کلید به فیلتر بلوم
		}
	}
}
