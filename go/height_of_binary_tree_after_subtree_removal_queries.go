package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	heights := make([]int, 100001)
	depth := 0

	var walk func(*TreeNode, int)
	walk = func(n *TreeNode, d int) {
		if n == nil {
			return
		}
		heights[n.Val] = depth
		depth = max(depth, d)
		walk(n.Left, d+1)
		walk(n.Right, d+1)
	}
	walk(root, 0)

	depth = 0
	walk = func(n *TreeNode, d int) {
		if n == nil {
			return
		}
		heights[n.Val] = max(heights[n.Val], depth)
		depth = max(depth, d)
		walk(n.Right, d+1)
		walk(n.Left, d+1)
	}
	walk(root, 0)

	ret := make([]int, len(queries))
	for i, q := range queries {
		ret[i] = heights[q]
	}
	return ret
}
