package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, delete []int) []*TreeNode {
	del := [1001]bool{}
	for _, v := range delete {
		del[v] = true
	}
	roots := []*TreeNode{}
	var walk func(*TreeNode, **TreeNode, bool)
	walk = func(n *TreeNode, np **TreeNode, root bool) {
		if n == nil {
			return
		}
		if del[n.Val] {
			if np != nil {
				*np = nil
			}
			walk(n.Left, nil, true)
			walk(n.Right, nil, true)
		} else {
			if root {
				roots = append(roots, n)
			}
			walk(n.Left, &n.Left, false)
			walk(n.Right, &n.Right, false)
		}
	}
	walk(root, nil, true)
	return roots
}
