package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	vals []int
}

func walk(t *TreeNode) []int {
	if t == nil {
		return nil
	}
	vals := []int{}
	vals = append(vals, walk(t.Left)...)
	vals = append(vals, t.Val)
	vals = append(vals, walk(t.Right)...)
	return vals
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{walk(root)}
}

func (i *BSTIterator) Next() int {
	x := i.vals[0]
	i.vals = i.vals[1:]
	return x
}

func (i *BSTIterator) HasNext() bool {
	return len(i.vals) > 0
}
