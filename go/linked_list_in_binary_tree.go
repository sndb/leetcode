package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	return helper(head, root) || root != nil && (isSubPath(head, root.Left) || isSubPath(head, root.Right))
}

func helper(head *ListNode, root *TreeNode) bool {
	return head == nil || root != nil && head.Val == root.Val && (helper(head.Next, root.Left) || helper(head.Next, root.Right))
}
