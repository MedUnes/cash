package main

import (
	"fmt"
	"log"

	"github.com/medunes/cash/cache"
)

func main() {
	htmlFiles := map[string]string{
		"file1.txt": `
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
		`, "file2.txt": `
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.

		`, "file3.txt": `
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
			This is a text file used to test LRU cache performance.
		`,
	}
	c, err := cache.NewLRUCache[string, string](1000)
	if err != nil {
		log.Fatalf("error initializing cache (%s)", err.Error())
		return
	}
	for name, content := range htmlFiles {
		c.Put(name, content)
	}
	fmt.Printf("Cache entries:\n")
	for name := range htmlFiles {
		content, ok := c.Get(name)
		if !ok {
			fmt.Printf("Skipping entry %s:\n", name)

			continue
		}

		fmt.Printf("\t%s: (%d bytes)\n", name, len(content))
	}
}
