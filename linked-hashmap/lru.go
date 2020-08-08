package main

type LRU interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Len() int
	Cap() int
	Del(key string) (string, bool)
	OrderedKeys() []string
}

type node struct {
	k    string
	v    string
	prev *node
	next *node
}

type LRUCache struct {
	list     *node
	cache    map[string]*node
	capacity int
}

func NewLRU(c int) LRU {
	return &LRUCache{
		list:     nil,
		cache:    map[string]*node{},
		capacity: c,
	}
}

// this assumes that n is already in the list
func (c *LRUCache) moveFront(n *node) {
	//is the node already the head?
	// nothing to do
	if c.list == n {
		return
	}

	// is the node already the tail?
	if n.next == nil {
		newTail := n.prev
		newTail.next = nil

		head := c.list
		head.prev = n
		n.prev = nil
		n.next = head
		c.list = n

		return
	}

	// is it in the middle
	// this removes it
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev
	// add it to the front
	head := c.list
	head.prev = n
	n.prev = nil
	n.next = head
	c.list = n
}

func (c *LRUCache) Get(key string) (string, bool) {
	n, ok := c.cache[key]
	if !ok {
		return "", false
	}

	// its there, so move it to the front
	c.moveFront(n)
	return n.v, true
}

func (c *LRUCache) Set(key, value string) {
	n, ok := c.cache[key]
	if !ok {
		// add it to the front
		newNode := &node{key, value, nil, c.list}
		c.cache[key] = newNode

		if c.list != nil {
			c.list.prev = newNode
		}

		c.list = newNode
		return
	}

	n.v = value
	c.moveFront(n)
	return
}

func (c *LRUCache) Del(key string) (string, bool) {
	n, ok := c.cache[key]
	if !ok {
		return "", false
	}
	delete(c.cache, key)

	// head
	if n == c.list {
		c.list = c.list.next
		return n.v, true
	}

	// tail
	if n.next == nil {
		n.prev.next = nil
		return n.v, true
	}

	// middle
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev
	return n.v, true
}

func (c *LRUCache) Len() int {
	return len(c.cache)
}

func (c *LRUCache) Cap() int {
	return c.capacity
}

func (c *LRUCache) OrderedKeys() []string {
	keys := make([]string, 0, c.Len())

	for n := c.list; n != nil; n = n.next {
		keys = append(keys, n.k)
	}
	return keys
}
