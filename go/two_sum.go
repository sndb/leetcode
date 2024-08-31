package main

func twoSum(nums []int, target int) []int {
	a := map[int]int{}
	for i, n := range nums {
		if j, ok := a[target-n]; ok {
			return []int{i, j}
		}
		a[n] = i
	}
	return nil
}
