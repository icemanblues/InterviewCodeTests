package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLRU(t *testing.T) {
	lru := NewLRU(10)
	assert.NotNil(t, lru)
	assert.Equal(t, 0, lru.Len())
}

func TestSet(t *testing.T) {
	lru := NewLRU(10)
	assert.Equal(t, 0, lru.Len())

	lru.Set("a", "a")
	assert.Equal(t, 1, lru.Len())

	lru.Set("b", "b")
	assert.Equal(t, 2, lru.Len())

	lru.Set("a", "aa")
	assert.Equal(t, 2, lru.Len())
}

func TestGet(t *testing.T) {
	lru := NewLRU(10)

	_, ok := lru.Get("a")
	assert.False(t, ok)

	lru.Set("a", "a")
	act, ok := lru.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "a", act)

	lru.Set("b", "b")
	act, ok = lru.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "b", act)

	lru.Set("a", "aa")
	act, ok = lru.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "aa", act)
}

func TestOrderedKeys(t *testing.T) {
	lru := NewLRU(10)
	assert.Equal(t, 0, lru.Len())
	assert.Zero(t, len(lru.OrderedKeys()))
	assert.Equal(t, []string{}, lru.OrderedKeys())

	lru.Set("a", "a")
	assert.Equal(t, []string{"a"}, lru.OrderedKeys())
	lru.Set("b", "b")
	assert.Equal(t, []string{"b", "a"}, lru.OrderedKeys())
	lru.Set("c", "c")
	assert.Equal(t, []string{"c", "b", "a"}, lru.OrderedKeys())

	lru.Set("a", "aa")
	assert.Equal(t, []string{"a", "c", "b"}, lru.OrderedKeys())

	lru.Del("c")
	assert.Equal(t, []string{"a", "b"}, lru.OrderedKeys())
}

func cache3() LRU {
	l := NewLRU(10)
	l.Set("a", "a")
	l.Set("b", "b")
	l.Set("c", "c")
	return l
}

func TestDel(t *testing.T) {
	tests := []struct {
		name  string
		keys  []string
		del   string
		value string
		ok    bool
		order []string
	}{
		{
			name:  "empty",
			keys:  []string{},
			del:   "a",
			value: "",
			ok:    false,
			order: []string{},
		},
		{
			name:  "single",
			keys:  []string{"a"},
			del:   "a",
			value: "a",
			ok:    true,
			order: []string{},
		},
		{
			name:  "head",
			keys:  []string{"a", "b", "c"},
			del:   "a",
			value: "a",
			ok:    true,
			order: []string{"c", "b"},
		},
		{
			name:  "middle",
			keys:  []string{"a", "b", "c"},
			del:   "b",
			value: "b",
			ok:    true,
			order: []string{"c", "a"},
		},
		{
			name:  "tail",
			keys:  []string{"a", "b", "c"},
			del:   "c",
			value: "c",
			ok:    true,
			order: []string{"b", "a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lru := NewLRU(10)
			for _, k := range test.keys {
				lru.Set(k, k)
			}

			actual, ok := lru.Del(test.del)
			assert.Equal(t, test.ok, ok)
			assert.Equal(t, test.value, actual)
			assert.Equal(t, test.order, lru.OrderedKeys())
		})
	}
}
