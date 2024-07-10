package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var walk func(*TreeNode, int, []int)
	walk = func(n *TreeNode, a int, v []int) {
		if n == nil {
			return
		}
		a += n.Val
		v = append(v, n.Val)
		if n.Left == nil && n.Right == nil && a == targetSum {
			w := make([]int, len(v))
			copy(w, v)
			result = append(result, w)
		}
		walk(n.Left, a, v)
		walk(n.Right, a, v)
	}
	walk(root, 0, nil)
	return result
}
