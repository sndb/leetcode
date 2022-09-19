package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderN(root *TreeNode, r map[int][]int, n int) {
	if root == nil {
		return
	}
	r[n] = append(r[n], root.Val)
	levelOrderN(root.Left, r, n+1)
	levelOrderN(root.Right, r, n+1)
}

func levelOrder(root *TreeNode) [][]int {
	r := map[int][]int{}
	levelOrderN(root, r, 0)
	l := make([][]int, len(r))
	for k, v := range r {
		l[k] = v
	}
	return l
}

func main() {
	tree := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	fmt.Printf("%#v\n", levelOrder(tree))
}
