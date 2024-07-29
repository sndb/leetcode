package main

func numTeams(rating []int) int {
	a := 0
	n := len(rating)
	for i := 0; i < n; i++ {
		l := 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				l++
			}
		}
		r := 0
		for j := i + 1; j < n; j++ {
			if rating[j] > rating[i] {
				r++
			}
		}
		a += l*r + (i-l)*(n-i-1-r)
	}
	return a
}
