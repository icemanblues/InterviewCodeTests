package main

import "fmt"

type node struct {
	k, v interface{}
	n, p *node
}

type LinkedHashmap struct {
	list *node
	hash map[interface{}]*node
}

func New() *LinkedHashmap {
	return &LinkedHashmap{hash: make(map[interface{}]*node)}
}

func (l *LinkedHashmap) Put(k, v interface{}) {
	// check if it already exists. if not, create it and add to map
	n, ok := l.hash[k]
	if !ok {
		n = &node{k: k, v: v}
	} else {
		n.k, n.v = k, v
	}

	l.hash[k] = n
	l.touched(n)
}

// touched helper function to move this n to the front of the list
func (l *LinkedHashmap) touched(n *node) {
	// if the head is empty, you are the head
	if l.list == nil {
		l.list = n
	}

	// this node is already the head of the list, nothing to do here
	if n == l.list {
		return
	}

	// now move the node to the front of the list
	// remove the node from the list
	prev, next := n.p, n.n
	if prev != nil {
		prev.n = next
	}
	if next != nil {
		next.p = prev
	}
	// insert the node to the head of the list
	n.p = nil
	n.n = l.list
	l.list.p = n
	// update the head pointer
	l.list = n
}

func (l *LinkedHashmap) del(n *node) {
	// this node is already the head of the list, just increment the list
	if n == l.list {
		l.list = l.list.n
		return
	}
	// not the head of the list, but still within the list
	prev, next := n.p, n.n
	if prev != nil {
		prev.n = next
	}
	if next != nil {
		next.p = prev
	}
}

func (l *LinkedHashmap) Get(k interface{}) (interface{}, bool) {
	n, ok := l.hash[k]
	if ok {
		l.touched(n)
		return n.v, true
	}
	return nil, false
}

func (l *LinkedHashmap) Remove(k interface{}) (interface{}, bool) {
	n, ok := l.hash[k]
	if !ok {
		return nil, false
	}

	// found it, so remove the node from the list
	delete(l.hash, k)
	l.del(n)
	return n.v, true
}

func (l *LinkedHashmap) Len() int {
	return len(l.hash)
}

func (l *LinkedHashmap) Youngest() (k, v interface{}) {
	if l.list == nil {
		return nil, nil
	}
	return l.list.k, l.list.v
}

func (l *LinkedHashmap) Oldest() (k, v interface{}) {
	if l.list == nil {
		return nil, nil
	}

	// FIXME: This is not constant time
	// FIXME: It would need another node pointer as the tail of the list
	n := l.list
	for n.n != nil {
		n = n.n
	}
	return n.k, n.v
}

func main() {
	fmt.Println("Linked Hashmap")

	lm := New()
	fmt.Printf("newly created linked hashmap: %v\n", lm)
	fmt.Printf("newly created with size: %v\n", lm.Len())

	lm.Put("hi", "hola")
	fmt.Println("added hi=>hola")

	fmt.Printf("added one element: %v\n", lm.Len())
	yK, yV := lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV := lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	lm.Put("world", "mundial")
	fmt.Println("added world=>mundial")
	lm.Put("bye", "adios")
	fmt.Println("added bye=>adios")

	fmt.Printf("added two more elements: %v\n", lm.Len())
	yK, yV = lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV = lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	value, ok := lm.Get("does not exist")
	fmt.Printf("Get key that doesnt exist: %v %v\n", value, ok)

	value, ok = lm.Get("bye")
	fmt.Printf("Get bye: %v %v\n", value, ok)
	yK, yV = lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV = lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	value, ok = lm.Get("world")
	fmt.Printf("Get world: %v %v\n", value, ok)
	yK, yV = lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV = lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	// delete a key that does not exist
	value, ok = lm.Remove("does not exist")
	fmt.Printf("Remove key that doesnt exist: %v %v\n", value, ok)
	yK, yV = lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV = lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	// delete a key that does exist
	value, ok = lm.Remove("hi")
	fmt.Printf("Remove hi: %v %v\n", value, ok)
	yK, yV = lm.Youngest()
	fmt.Printf("youngest pair: %v %v\n", yK, yV)
	oldK, oldV = lm.Oldest()
	fmt.Printf("oldest pair: %v %v\n", oldK, oldV)

	fmt.Println("exiting...")
}
