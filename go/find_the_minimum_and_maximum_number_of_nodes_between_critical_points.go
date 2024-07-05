package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	points := []int{}
	for i, n := 0, head; n.Next.Next != nil; i, n = i+1, n.Next {
		p, c, s := n.Val, n.Next.Val, n.Next.Next.Val
		if c > p && c > s || c < p && c < s {
			points = append(points, i)
		}
	}
	if len(points) < 2 {
		return []int{-1, -1}
	}
	r1 := points[1] - points[0]
	r2 := points[len(points)-1] - points[0]
	for i := 2; i < len(points); i++ {
		r1 = min(r1, points[i]-points[i-1])
	}
	return []int{r1, r2}
}
