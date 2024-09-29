package main

type AllOne struct {
	items map[string]*node
	root  *node
}

func Constructor() AllOne {
	items := map[string]*node{}

	root := newNode(1)
	root.next = root
	root.prev = root

	return AllOne{items, root}
}

func (a *AllOne) Inc(key string) {
	n, ok := a.items[key]
	if !ok {
		a.add(a.root, key)
		return
	}

	if n.next.index == n.index+1 {
		a.add(n.next, key)
	} else {
		m := newNode(n.index + 1)
		a.add(m, key)
		m.prev = n
		m.next = n.next
		n.next.prev = m
		n.next = m
	}
	a.cleanup(n, key)
}

func (a *AllOne) Dec(key string) {
	n := a.items[key]

	if n.index == 1 {
		delete(a.items, key)
	} else if n.prev.index == n.index-1 {
		a.add(n.prev, key)
	} else {
		m := newNode(n.index - 1)
		a.add(m, key)
		m.next = n
		m.prev = n.prev
		n.prev.next = m
		n.prev = m
	}
	a.cleanup(n, key)
}

func (a *AllOne) GetMaxKey() string {
	for k := range a.root.prev.strings {
		return k
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	for k := range a.root.strings {
		return k
	}
	for k := range a.root.next.strings {
		return k
	}
	return ""
}

func (a *AllOne) cleanup(n *node, key string) {
	delete(n.strings, key)
	if n != a.root && len(n.strings) == 0 {
		n.prev.next, n.next.prev = n.next, n.prev
	}
}

func (a *AllOne) add(n *node, key string) {
	a.items[key] = n
	n.strings[key] = true
}

type node struct {
	next, prev *node
	strings    map[string]bool
	index      int
}

func newNode(index int) *node {
	return &node{strings: map[string]bool{}, index: index}
}
