package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{}
	evenHead := &ListNode{}
	oddTail := oddHead
	evenTail := evenHead
	flag := false
	for head != nil {
		next := head.Next
		if flag {
			evenTail.Next = head
			evenTail = head
		} else {
			oddTail.Next = head
			oddTail = head
		}
		flag = !flag
		head = next
	}
	oddTail.Next = evenHead.Next
	evenTail.Next = nil
	return oddHead.Next
}
