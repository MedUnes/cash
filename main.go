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
