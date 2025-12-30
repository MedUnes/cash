package cache

type node[K comparable, V any] struct {
	next, prev *node[K, V]
	item       item[K, V]
}
type linkedList[K comparable, V any] struct {
	head, tail *node[K, V]
}

func newLinkedList[K comparable, V any]() linkedList[K, V] {
	return linkedList[K, V]{}
}

func (l *linkedList[K, V]) moveToFront(n *node[K, V]) {
	if l.head == n {
		return
	}
	l.remove(n)
	l.pushFront(n)
	return
}

func (l *linkedList[K, V]) remove(n *node[K, V]) {
	if n == nil {
		return
	}
	if l.head == n {
		l.head = n.next
	}
	if l.tail == n {
		l.tail = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

func (l *linkedList[K, V]) pushFront(n *node[K, V]) {
	n.next = l.head
	n.prev = nil
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
}
