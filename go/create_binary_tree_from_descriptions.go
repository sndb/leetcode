package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	children := map[*TreeNode]bool{}
	nodes := map[int]*TreeNode{}
	for _, d := range descriptions {
		parent, ok := nodes[d[0]]
		if !ok {
			parent = &TreeNode{Val: d[0]}
			nodes[d[0]] = parent
		}
		child, ok := nodes[d[1]]
		if !ok {
			child = &TreeNode{Val: d[1]}
			nodes[d[1]] = child
		}
		children[child] = true
		if d[2] == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}
	}
	for _, n := range nodes {
		if !children[n] {
			return n
		}
	}
	return nil
}
