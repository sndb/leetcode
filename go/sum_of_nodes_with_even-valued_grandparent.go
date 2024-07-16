package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum int
	var walk func(*TreeNode, *TreeNode, *TreeNode)
	walk = func(gp, p, n *TreeNode) {
		if n == nil {
			return
		}
		if gp != nil && gp.Val%2 == 0 {
			sum += n.Val
		}
		walk(p, n, n.Left)
		walk(p, n, n.Right)
	}
	walk(nil, nil, root)
	return sum
}
