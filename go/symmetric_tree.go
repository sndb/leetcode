package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorEqual(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && mirrorEqual(left.Right, right.Left) && mirrorEqual(left.Left, right.Right)
}

func isSymmetric(root *TreeNode) bool {
	return mirrorEqual(root.Left, root.Right)
}
