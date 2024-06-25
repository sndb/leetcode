package main

func bstToGst(root *TreeNode) *TreeNode {
	acc := 0
	var convert func(*TreeNode)
	convert = func(t *TreeNode) {
		if t == nil {
			return
		}
		convert(t.Right)
		acc += t.Val
		t.Val = acc
		convert(t.Left)
	}
	convert(root)
	return root
}
