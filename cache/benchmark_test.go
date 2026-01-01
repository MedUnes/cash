package cache

import (
	"testing"
)

func BenchmarkPut(b *testing.B) {
	c, _ := NewLRUCache[int, int](10000)
	for i := range b.N {
		c.Put(i, i)
	}
}
