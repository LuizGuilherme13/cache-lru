package main

import (
	"fmt"
)

func main() {
	cache := New(3)

	for i, key := range []string{"A", "B", "C"} {
		cache.Set(key, i)
	}

	for _, key := range []string{"C", "A", "B"} {
		cache.Get(key)
	}

	for _, item := range cache.items {
		fmt.Printf("%s : %v\n", item.key, item.access)
	}
	fmt.Println("==================================")

	cache.Set("D", 40)

	for _, item := range cache.items {
		fmt.Printf("%s : %v\n", item.key, item.access)
	}
	fmt.Println("==================================")

	cache.Set("E", 50)

	for _, item := range cache.items {
		fmt.Printf("%s : %v\n", item.key, item.access)
	}
	fmt.Println("==================================")

	cache.Get("E")

	for _, item := range cache.items {
		fmt.Printf("%s : %v\n", item.key, item.access)
	}
}
