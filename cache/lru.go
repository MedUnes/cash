package cache

import (
	"fmt"
	"sync"
)

type item[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity uint64
	m        map[K]*node[K, V]
	mu       sync.Mutex
	records  linkedList[K, V]
}

func NewLRUCache[K comparable, V any](c uint64) (*LRUCache[K, V], error) {
	if c == 0 {
		return nil, fmt.Errorf("capacity can not be 0")

	}
	l := &LRUCache[K, V]{}
	l.m = make(map[K]*node[K, V])
	l.capacity = c
	l.records = newLinkedList[K, V]()
	return l, nil
}
func (l *LRUCache[K, V]) Put(key K, value V) {
	l.mu.Lock()
	defer l.mu.Unlock()
	n, ok := l.m[key]
	if ok {
		n.item.value = value
		l.records.moveToFront(n)
		return
	}

	if uint64(len(l.m)) == l.capacity {
		delete(l.m, l.records.tail.item.key)
		l.records.remove(l.records.tail)
	}
	node := &node[K, V]{nil, nil, item[K, V]{key, value}}
	l.records.moveToFront(node)
	l.m[key] = node
}
func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	node, ok := l.m[key]
	if ok {
		l.records.moveToFront(node)
		return node.item.value, true
	}
	var zero V
	return zero, false
}
