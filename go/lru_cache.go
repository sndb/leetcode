package main

import "fmt"

type LRUCache struct {
	items map[int]*node
	root  *node
	cap   int
}

func Constructor(cap int) LRUCache {
	keys := make(map[int]*node)

	root := &node{}
	root.next = root
	root.prev = root

	return LRUCache{keys, root, cap}
}

func (l *LRUCache) Get(key int) int {
	n, ok := l.items[key]
	if !ok {
		return -1
	}
	l.update(n)
	return n.val
}

func (l *LRUCache) Put(key, val int) {
	n, ok := l.items[key]
	if ok {
		n.val = val
	} else {
		n = &node{key: key, val: val}
		n.next = n
		n.prev = n
	}
	l.items[key] = n
	l.update(n)
	l.trim()
}

func (l *LRUCache) update(n *node) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev, n.next = l.root.prev, l.root
	l.root.prev.next, l.root.prev = n, n
}

func (l *LRUCache) trim() {
	if len(l.items) > l.cap {
		delete(l.items, l.root.next.key)
		l.root.next = l.root.next.next
		l.root.next.prev = l.root
	}
}

type node struct {
	prev, next *node
	key, val   int
}

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println(l.Get(2))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
