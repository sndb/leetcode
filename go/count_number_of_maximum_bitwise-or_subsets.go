package main

func countMaxOrSubsets(nums []int) int {
	or := 0
	ret := 0
	seen := []int{}
	for _, n := range nums {
		add := []int{n}
		for _, m := range seen {
			add = append(add, n|m)
		}
		seen = append(seen, add...)
		or |= n
	}
	for _, n := range seen {
		if n == or {
			ret++
		}
	}
	return ret
}
