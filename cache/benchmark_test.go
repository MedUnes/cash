package cache

import (
	"fmt"
	"testing"
)

func preparePut[K int | string, V any](cacheSize int, keyGen func(int) K, valGen func(int) V) map[K]V {
	m := make(map[K]V, cacheSize)
	for i := range cacheSize {
		m[keyGen(i)] = valGen(i)
	}

	return m
}

func prepareGet[K comparable, V any](cacheSize int, keyGen func(int) K, valGen func(int) V) *LRUCache[K, V] {
	s := uint64(cacheSize)
	c, _ := NewLRUCache[K, V](s)
	keys := make([]K, cacheSize)
	for i := range cacheSize {
		keys[i] = keyGen(i)
	}

	for i := range cacheSize {
		c.Put(keys[i], valGen(i))
	}
	return c
}

func BenchmarkPutIntKey(b *testing.B) {
	cacheSize := 100
	m := preparePut[int, int](cacheSize, func(i int) int { return i }, func(i int) int { return i })
	c, _ := NewLRUCache[int, int](uint64(cacheSize))
	b.ResetTimer()
	for i := range b.N {
		c.Put(i%cacheSize, m[i%cacheSize])
	}
}

func BenchmarkPutStringKey(b *testing.B) {
	cacheSize := 100
	m := preparePut[int, string](cacheSize, func(i int) int { return i }, func(i int) string { return fmt.Sprintf("key-%d", i) })
	c, _ := NewLRUCache[string, int](uint64(cacheSize))
	b.ResetTimer()
	for i := range b.N {
		c.Put(m[i%cacheSize], i)
	}
}

func BenchmarkGetIntKey(b *testing.B) {
	cacheSize := 100
	c := prepareGet[int, int](cacheSize, func(i int) int { return i }, func(i int) int { return i })
	b.ResetTimer()
	for i := range b.N {
		_, _ = c.Get(i % cacheSize)
	}
}

func BenchmarkGetStringKey(b *testing.B) {
	cacheSize := 100
	m := preparePut[int, string](cacheSize, func(i int) int { return i }, func(i int) string { return fmt.Sprintf("key-%d", i) })
	c := prepareGet[string, int](cacheSize, func(i int) string { return fmt.Sprintf("key-%d", i) }, func(i int) int { return i })
	b.ResetTimer()
	for i := range b.N {
		_, _ = c.Get(m[i%cacheSize])
	}
}

func BenchmarkEviction(b *testing.B) {
	cacheSize := 1000
	c, _ := NewLRUCache[int, int](uint64(cacheSize))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Put(i, i)
	}
}

func BenchmarkParallel(b *testing.B) {
	cacheSize := 1000
	c, _ := NewLRUCache[int, int](uint64(cacheSize))

	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			counter++
			// 50% Writes, 50% Reads: a common chaotic pattern
			if counter%2 == 0 {
				c.Put(counter, counter)
			} else {
				// Read a "hot" key (10%) to force contention
				c.Get(counter % 10)
			}
		}
	})
}
