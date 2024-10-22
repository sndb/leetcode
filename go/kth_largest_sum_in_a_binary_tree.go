package main

import (
	"maps"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levels := map[int]int64{}

	var walk func(*TreeNode, int)
	walk = func(n *TreeNode, d int) {
		if n != nil {
			levels[d] += int64(n.Val)
			walk(n.Left, d+1)
			walk(n.Right, d+1)
		}
	}
	walk(root, 0)

	sums := slices.Sorted(maps.Values(levels))
	i := len(sums) - k
	if i < 0 {
		return -1
	}
	return sums[i]
}
