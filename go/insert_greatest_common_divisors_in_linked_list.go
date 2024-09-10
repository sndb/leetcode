package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	prev, curr := head, head.Next
	for curr != nil {
		v1, v2 := prev.Val, curr.Val
		for v2 != 0 {
			v1, v2 = v2, v1%v2
		}
		prev.Next = &ListNode{v1, curr}
		prev, curr = curr, curr.Next
	}
	return head
}
