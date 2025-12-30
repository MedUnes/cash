# cash

[![Go Report Card](https://goreportcard.com/badge/github.com/medunes/cash)](https://goreportcard.com/report/github.com/medunes/cash)
[![Go Reference](https://pkg.go.dev/badge/github.com/medunes/cash.svg)](https://pkg.go.dev/github.com/medunes/cash)
[![Tests](https://github.com/medunes/cash/actions/workflows/tests.yml/badge.svg)](https://github.com/medunes/cash/actions/workflows/tests.yml)

* An implementation of the [LRU Cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU) data structure.
* Read/Write are in O(1) time complexity.

## Example usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/medunes/cash/cache"
)

func main() {
	c, err := cache.NewLRUCache[string, []byte](1000)
	if err != nil {
		log.Fatalf("error initializing cache (%s)", err.Error())
		return
	}
	htmlFiles, err := os.ReadDir("test/data")
	if err != nil {
		log.Fatalf("error loading html files (%s)", err.Error())
		return
	}
	for _, f := range htmlFiles {
		content, err := os.ReadFile("test/data/" + f.Name())
		if err != nil {
			log.Fatalf("error reading html file (%s)", err.Error())
			return
		}
		c.Put(f.Name(), content)
	}
	fmt.Printf("Cache entries:\n")
	for _, f := range htmlFiles {
		name := f.Name()
		content, ok := c.Get(name)
		if !ok {
			fmt.Printf("Skipping entry %s:\n", name)

			continue
		}

		fmt.Printf("\t%s: (%d Kb)\n", name, len(string(content))/1024)

	}
}
```

This should show something like:
```bash
go run main.go
Cache entries:
        ipsum_135kb.html: (134 Kb)
        ipsum_145kb.html: (144 Kb)
        ipsum_14kb.html: (13 Kb)

```