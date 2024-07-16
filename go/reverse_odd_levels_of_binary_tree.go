package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	odd := false
	queue := []*TreeNode{root}
	for queue[0] != nil {
		if odd {
			for i, n := 0, len(queue); i < n/2; i++ {
				a := queue[i]
				b := queue[n-i-1]
				a.Val, b.Val = b.Val, a.Val
			}
		}
		odd = !odd
		for n := len(queue); n > 0; n-- {
			car := queue[0]
			queue = queue[1:]
			queue = append(queue, car.Left, car.Right)
		}
	}
	return root
}
