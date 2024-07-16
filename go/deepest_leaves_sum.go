package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	lim := 0
	sum := 0
	var walk func(*TreeNode, int)
	walk = func(n *TreeNode, d int) {
		if n == nil {
			return
		}
		if d > lim {
			lim = d
			sum = 0
		}
		if d == lim {
			sum += n.Val
		}
		walk(n.Left, d+1)
		walk(n.Right, d+1)
	}
	walk(root, 0)
	return sum
}
