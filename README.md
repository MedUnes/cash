# cash
[![Release](https://github.com/MedUnes/cash/actions/workflows/release.yaml/badge.svg)](https://github.com/MedUnes/cash/actions/workflows/release.yaml)
[![Tests](https://github.com/MedUnes/cash/actions/workflows/tests.yaml/badge.svg)](https://github.com/MedUnes/cash/actions/workflows/tests.yaml)
[![CodeQL](https://github.com/MedUnes/cash/actions/workflows/codeql.yaml/badge.svg)](https://github.com/MedUnes/cash/actions/workflows/codeql.yaml)
[![Trivy Security Scan](https://github.com/MedUnes/cash/actions/workflows/trivy.yaml/badge.svg)](https://github.com/MedUnes/cash/actions/workflows/trivy.yaml)
[![codecov](https://codecov.io/gh/medunes/cash/branch/master/graph/badge.svg)](https://codecov.io/gh/medunes/cash)
[![Go Report Card](https://goreportcard.com/badge/github.com/medunes/cash)](https://goreportcard.com/report/github.com/medunes/cash)
[![Go Reference](https://pkg.go.dev/badge/github.com/medunes/cash.svg)](https://pkg.go.dev/github.com/medunes/cash)

**cash** is a thread-safe, generic LRU (Least Recently Used) cache implementation for Go. 

- It is designed for high-throughput systems where Garbage Collection (GC) pauses are unacceptable. 
- It achieves **O(1)** time complexity for all operations and **zero heap allocations** (`0 allocs/op`) during standard read/write cycles.

## Features

* Zero GC Overhead: Optimized to generate **0 allocations** for `Put` (updates) and `Get` operations.
* Thread-Safe: Built-in `sync.Mutex` protection ensures safety for concurrent access.
* Generic: leveraging Go 1.18+ generics (`[K comparable, V any]`) for type safety without reflection.
* Production Ready: 100% test coverage, validated with automated security scans (Trivy/CodeQL) and race detectors.

## Performance

Benchmarks were run on a 12th Gen Intel i5-12500H. 

The library maintains **0 allocs/op** even under high concurrency, ensuring that the cache does not contribute to GC pressure in latency-sensitive applications.

| Benchmark | Iterations | ns/op | B/op | allocs/op | Note |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Put (Int Key)** | 31M | **38.09 ns** | 0 | **0** | No allocation overhead |
| **Put (String Key)** | 27M | 42.50 ns | 0 | **0** | No allocation overhead |
| **Get (Int Key)** | 36M | **32.40 ns** | 0 | **0** | Instant retrieval |
| **Parallel R/W** | 4M | 303.1 ns | 14 | **0** | Thread-safe, no GC spikes |
| **Eviction** | 7M | 161.9 ns | 32 | 1 | New node creation |

## Installation

```bash
go get https://github.com/medunes/cash

```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/medunes/cash/cache"
)

func main() {
	// Initialize a cache with capacity 1000
	// Key: string, Value: []byte
	c, err := cache.NewLRUCache[string, []byte](1000)
	if err != nil {
		log.Fatal(err)
	}

	// O(1) Write
	c.Put("user:123", []byte("data"))

	// O(1) Read
	if val, ok := c.Get("user:123"); ok {
		fmt.Printf("Found: %s\n", string(val))
	}
}

```

## Testing

The project maintains rigorous quality standards with 19 unit tests covering edge cases (empty capacity, idempotency, eviction logic).

```bash
$ make test
✓ Cache (0.00s)
✓ Cache cache eviction (0.00s)
✓ Linked list remove (0.00s)
...
DONE 19 tests in 0.323s

```