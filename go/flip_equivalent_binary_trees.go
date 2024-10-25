package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1, root2 *TreeNode) bool {
	var walk func(*TreeNode, [][2]int)
	walk = func(n *TreeNode, h [][2]int) {
		if n != nil {
			h[n.Val] = [2]int{-1, -1}
			if n.Left != nil {
				h[n.Val][0] = n.Left.Val + 1
			}
			if n.Right != nil {
				h[n.Val][1] = n.Right.Val + 1
			}
			if h[n.Val][1] < h[n.Val][0] {
				h[n.Val][0], h[n.Val][1] = h[n.Val][1], h[n.Val][0]
			}
			walk(n.Left, h)
			walk(n.Right, h)
		}
	}
	h1 := make([][2]int, 100)
	h2 := make([][2]int, 100)
	walk(root1, h1)
	walk(root2, h2)
	return slices.Equal(h1, h2)
}
