package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	index := map[*Node]*Node{}
	var tail *Node
	for node := head; node != nil; node = node.Next {
		new := &Node{Val: node.Val}
		if tail != nil {
			tail.Next = new
		}
		tail = new
		index[node] = new
	}
	for node := head; node != nil; node = node.Next {
		index[node].Random = index[node.Random]
	}
	return index[head]
}
