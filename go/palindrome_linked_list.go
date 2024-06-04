package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Val)
	}
	for l, r := 0, len(vals)-1; l < r; l, r = l+1, r-1 {
		if vals[l] != vals[r] {
			return false
		}
	}
	return true
}
