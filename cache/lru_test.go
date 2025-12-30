package cache

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLRUCache(t *testing.T) {
	t.Run("zero capacity does not build", func(t *testing.T) {
		c, err := NewLRUCache[string, int](0)
		assert.Nil(t, c)
		assert.Error(t, err)
	})
	t.Run("least minimal capacity builds", func(t *testing.T) {
		c, err := NewLRUCache[string, int](1)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), c.capacity)
		assert.NotNil(t, c.capacity)
	})

	t.Run("just least than max capacity builds", func(t *testing.T) {
		c, err := NewLRUCache[string, int](uint64(math.MaxUint64) - 1)
		assert.Nil(t, err)
		assert.Equal(t, uint64(uint64(math.MaxUint64)-1), c.capacity)
	})
	t.Run("max capacity builds", func(t *testing.T) {
		c, err := NewLRUCache[string, int](uint64(math.MaxUint64))
		assert.Nil(t, err)
		assert.Equal(t, uint64(math.MaxUint64), c.capacity)
	})
}

func TestCache(t *testing.T) {
	type Put struct {
		key string
		val int
	}
	type Expected struct {
		ok  bool
		val int
	}
	type Get struct {
		key      string
		expected Expected
	}
	type Op struct {
		put []Put
		get []Get
	}
	testData := []struct {
		name     string
		capacity uint64
		ops      []Op
	}{
		{
			name:     "capacity 1 single put",
			capacity: 1,
			ops: []Op{
				{
					[]Put{{"a", 1}},
					[]Get{{"a", Expected{true, 1}}},
				},
			},
		},
		{
			name:     "Idempotent put",
			capacity: 10,
			ops: []Op{
				{
					[]Put{{"a", 1}},
					[]Get{{"a", Expected{true, 1}}},
				},
				{
					[]Put{{"a", 1}},
					[]Get{{"a", Expected{true, 1}}},
				},
				{
					[]Put{{"a", 1}},
					[]Get{{"a", Expected{true, 1}}},
				},
			},
		},
		{
			name:     "Idempotent Get",
			capacity: 7,
			ops: []Op{
				{
					put: []Put{{"a", 1}},
					get: []Get{
						{"a", Expected{true, 1}},
						{"a", Expected{true, 1}},
						{"a", Expected{true, 1}},
					},
				},
			},
		},
		{
			name:     "Missing key",
			capacity: 7,
			ops: []Op{
				{
					[]Put{{"a", 1}, {"b", 2}, {"abc", 123}},
					[]Get{
						{"a", Expected{true, 1}},
						{"b", Expected{true, 2}},
						{"abc", Expected{true, 123}},
						{"x", Expected{false, 0}},
						{"ab", Expected{false, 0}},
						{"abx", Expected{false, 0}},
						{"abcd", Expected{false, 0}},
						{"xxx", Expected{false, 0}},
						{"xxxabc", Expected{false, 0}},
						{"xxxabcd", Expected{false, 0}},
					},
				},
			},
		},
		{
			name:     "Cache Eviction",
			capacity: 2,
			ops: []Op{
				{
					[]Put{{"a", 1}, {"b", 123}},
					[]Get{
						{"a", Expected{true, 1}},
						{"a", Expected{true, 1}},
						{"b", Expected{true, 123}},
					},
				},
				{
					[]Put{{"xxx", 100}},
					[]Get{
						{"xxx", Expected{true, 100}},
						{"a", Expected{false, 0}},
						{"b", Expected{true, 123}},
					},
				},
			},
		},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			l, err := NewLRUCache[string, int](d.capacity)
			require.Nil(t, err)
			assert.Equal(t, d.capacity, l.capacity)
			for _, op := range d.ops {
				for _, p := range op.put {
					l.Put(p.key, p.val)
				}
				for _, g := range op.get {
					v, ok := l.Get(g.key)
					assert.Equal(t, g.expected.ok, ok)
					assert.Equal(t, g.expected.val, v)
				}
			}
		})
	}
}
