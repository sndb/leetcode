package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	zero := head
	for curr := head.Next; curr.Next != nil; curr = curr.Next {
		zero.Val += curr.Val
		if curr.Val == 0 {
			zero.Next = curr
			zero = curr
		}
	}
	zero.Next = nil
	return head
}
