package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	del := map[int]bool{}
	for _, n := range nums {
		del[n] = true
	}
	for prev, curr := head, head.Next; curr != nil; curr = curr.Next {
		if del[curr.Val] {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}
	if del[head.Val] {
		head = head.Next
	}
	return head
}
