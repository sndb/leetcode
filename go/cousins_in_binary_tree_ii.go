package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	for curr != nil {
		sum := 0
		for _, n := range curr {
			sum += nodeValue(n.Left) + nodeValue(n.Right)
		}
		for _, n := range curr {
			pair := nodeValue(n.Left) + nodeValue(n.Right)
			if n.Left != nil {
				n.Left.Val = sum - pair
				next = append(next, n.Left)
			}
			if n.Right != nil {
				n.Right.Val = sum - pair
				next = append(next, n.Right)
			}
		}
		curr, next = next, nil
	}
	return root
}

func nodeValue(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return n.Val
}
