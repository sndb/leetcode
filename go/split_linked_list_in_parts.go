package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := make([]int, k)
	for i, curr := 0, head; curr != nil; i, curr = i+1, curr.Next {
		count[i%k]++
	}
	parts := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		parts[i] = head
		prev := head
		for j := 1; j < count[i]; j++ {
			prev = prev.Next
		}
		if prev != nil {
			head = prev.Next
			prev.Next = nil
		}
	}
	return parts
}
