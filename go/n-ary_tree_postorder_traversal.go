package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var vals []int
	for _, c := range root.Children {
		vals = append(vals, postorder(c)...)
	}
	vals = append(vals, root.Val)
	return vals
}
