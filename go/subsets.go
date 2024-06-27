package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	r := [][]int{}
	x := nums[0]
	for _, s := range subsets(nums[1:]) {
		r = append(r, s)
		r = append(r, append([]int{x}, s...))
	}
	return r
}
