package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return fmt.Sprintf("_")
	}
	return fmt.Sprintf("(%v %s %s)", t.Val, t.Left, t.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{nums[mid], sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])}
}

func main() {
	fmt.Printf("%s\n", sortedArrayToBST([]int{1, 2, 3, 4, 7, 8, 10, 50}))
}
