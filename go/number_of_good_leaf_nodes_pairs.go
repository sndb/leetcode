package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	pairs := 0
	paths := [][]byte{}
	var walk func(*TreeNode, []byte)
	walk = func(n *TreeNode, p []byte) {
		if n == nil {
			return
		}
		if n.Left == n.Right {
			for _, q := range paths {
				n := 0
				m := min(len(p), len(q))
				for n < m && p[n] == q[n] {
					n++
				}
				if len(p[n:])+len(q[n:]) <= distance {
					pairs++
				}
			}
			paths = append(paths, slices.Clone(p))
		}
		walk(n.Left, append(p, 0))
		walk(n.Right, append(p, 1))
	}
	walk(root, nil)
	return pairs
}
