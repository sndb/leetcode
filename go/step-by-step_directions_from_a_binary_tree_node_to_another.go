package main

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPath(node *TreeNode, value int, path []byte) []byte {
	if node == nil {
		return nil
	}
	if node.Val == value {
		return path
	}
	if r := getPath(node.Left, value, append(path, 'L')); r != nil {
		return r
	}
	return getPath(node.Right, value, append(path, 'R'))
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	p1 := getPath(root, startValue, []byte{})
	p2 := getPath(root, destValue, []byte{})
	for len(p1) > 0 && len(p2) > 0 && p1[0] == p2[0] {
		p1 = p1[1:]
		p2 = p2[1:]
	}
	return strings.Repeat("U", len(p1)) + string(p2)
}
