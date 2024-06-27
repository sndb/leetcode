package main

type set [21]int

func subsetsWithDup(nums []int) [][]int {
	sets := map[set]bool{{}: true}
	for _, n := range nums {
		add := []set{}
		for s := range sets {
			s[n+10]++
			add = append(add, s)
		}
		for _, a := range add {
			sets[a] = true
		}
	}
	r := [][]int{}
	for s := range sets {
		x := []int{}
		for i, n := range s {
			for range n {
				x = append(x, i-10)
			}
		}
		r = append(r, x)
	}
	return r
}
