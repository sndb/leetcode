package main

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	a := sortArray(nums[:len(nums)/2])
	b := sortArray(nums[len(nums)/2:])
	r := make([]int, 0, len(nums))
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			r = append(r, a[0])
			a = a[1:]
		} else {
			r = append(r, b[0])
			b = b[1:]
		}
	}
	return append(append(r, a...), b...)
}
