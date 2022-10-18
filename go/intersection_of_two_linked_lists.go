package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(headA, headB *ListNode) <-chan *ListNode {
	ch := make(chan *ListNode)
	go func() {
		for headA != nil || headB != nil {
			if headA != nil {
				ch <- headA
				headA = headA.Next
			}
			if headB != nil {
				ch <- headB
				headB = headB.Next
			}
		}
		close(ch)
	}()
	return ch
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	ch := merge(headA, headB)
	for n := range ch {
		if seen[n] {
			return n
		}
		seen[n] = true
	}
	return nil
}
