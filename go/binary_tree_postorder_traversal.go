package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var vals []int
	vals = append(vals, postorderTraversal(root.Left)...)
	vals = append(vals, postorderTraversal(root.Right)...)
	vals = append(vals, root.Val)
	return vals
}
