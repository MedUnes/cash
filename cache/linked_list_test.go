package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	t.Run("create empty", func(t *testing.T) {
		l := newLinkedList[string, int]()
		assert.NotNil(t, l)
		assert.Nil(t, l.head)
		assert.Nil(t, l.tail)
	})
}

func TestLinkedListAddFirst(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		l := newLinkedList[string, int]()
		nodes := []*node[string, int]{
			{nil, nil, item[string, int]{"a", 1}},
			{nil, nil, item[string, int]{"b", 2}},
			{nil, nil, item[string, int]{"c", 3}},
		}
		for _, node := range nodes {
			l.pushFront(node)
		}
		assert.Equal(t, l.head.item.key, "c")
		assert.Equal(t, l.head.item.value, 3)
		assert.Equal(t, l.tail.item.key, "a")
		assert.Equal(t, l.tail.item.value, 1)
		assert.Equal(t, l.head.next.item.key, "b")
		assert.Equal(t, l.head.next.item.value, 2)
	})
}

func TestLinkedListRemove(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		l := newLinkedList[string, int]()
		nodes := []*node[string, int]{
			{nil, nil, item[string, int]{"a", 1}},
			{nil, nil, item[string, int]{"b", 2}},
			{nil, nil, item[string, int]{"c", 3}},
		}
		for _, node := range nodes {
			l.pushFront(node)
		}
		assert.Equal(t, l.head.item.key, "c")
		assert.Equal(t, l.head.item.value, 3)
		assert.Equal(t, l.tail.item.key, "a")
		assert.Equal(t, l.tail.item.value, 1)
		assert.Equal(t, l.head.next.item.key, "b")
		assert.Equal(t, l.head.next.item.value, 2)

		l.remove(l.head)
		assert.Equal(t, l.head.item.key, "b")
		assert.Equal(t, l.head.item.value, 2)
		assert.Equal(t, l.tail.item.key, "a")
		assert.Equal(t, l.tail.item.value, 1)

		l.remove(l.tail)
		assert.Equal(t, l.head.item.key, "b")
		assert.Equal(t, l.head.item.value, 2)
		assert.Equal(t, l.tail.item.key, "b")
		assert.Equal(t, l.tail.item.value, 2)
	})
}

func TestLinkedListMoveFirst(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		l := newLinkedList[string, int]()
		nodes := []*node[string, int]{
			{nil, nil, item[string, int]{"a", 1}},
			{nil, nil, item[string, int]{"b", 2}},
			{nil, nil, item[string, int]{"c", 3}},
		}
		for _, node := range nodes {
			l.pushFront(node)
		}
		assert.Equal(t, l.head.item.key, "c")
		assert.Equal(t, l.head.item.value, 3)
		assert.Equal(t, l.tail.item.key, "a")
		assert.Equal(t, l.tail.item.value, 1)
		assert.Equal(t, l.head.next.item.key, "b")
		assert.Equal(t, l.head.next.item.value, 2)

		l.moveToFront(l.head)
		assert.Equal(t, l.head.item.key, "c")
		assert.Equal(t, l.head.item.value, 3)
		assert.Equal(t, l.tail.item.key, "a")
		assert.Equal(t, l.tail.item.value, 1)
		assert.Equal(t, l.head.next.item.key, "b")
		assert.Equal(t, l.head.next.item.value, 2)

		l.moveToFront(l.tail)
		assert.Equal(t, l.head.item.key, "a")
		assert.Equal(t, l.head.item.value, 1)
		assert.Equal(t, l.tail.item.key, "b")
		assert.Equal(t, l.tail.item.value, 2)
		assert.Equal(t, l.head.next.item.key, "c")
		assert.Equal(t, l.head.next.item.value, 3)

		l.moveToFront(l.head.next)
		assert.Equal(t, l.head.item.key, "c")
		assert.Equal(t, l.head.item.value, 3)
		assert.Equal(t, l.tail.item.key, "b")
		assert.Equal(t, l.tail.item.value, 2)
		assert.Equal(t, l.head.next.item.key, "a")
		assert.Equal(t, l.head.next.item.value, 1)
	})
}
