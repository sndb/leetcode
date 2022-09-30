package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode, acc map[string][]*TreeNode) string {
	if root == nil {
		return "."
	}
	s := fmt.Sprintf(
		"%v(%v %v)",
		root.Val,
		serialize(root.Left, acc),
		serialize(root.Right, acc),
	)
	acc[s] = append(acc[s], root)
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	acc := map[string][]*TreeNode{}
	serialize(root, acc)
	r := []*TreeNode{}
	for _, v := range acc {
		if len(v) > 1 {
			r = append(r, v[0])
		}
	}
	return r
}

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(serialize(tree, map[string][]*TreeNode{})) // 1(2(. .) .)
}
