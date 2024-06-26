package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var vals []int
	var walk func(*TreeNode)
	walk = func(t *TreeNode) {
		if t == nil {
			return
		}
		walk(t.Left)
		vals = append(vals, t.Val)
		walk(t.Right)
	}
	walk(root)

	var build func([]int) *TreeNode
	build = func(a []int) *TreeNode {
		if len(a) == 0 {
			return nil
		}
		i := len(a) / 2
		return &TreeNode{a[i], build(a[:i]), build(a[i+1:])}
	}
	tree := build(vals)
	return tree
}
