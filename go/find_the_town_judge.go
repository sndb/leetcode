package main

func findJudge(n int, trust [][]int) int {
	a := make([]int, n+1)
	for _, t := range trust {
		a[t[0]]--
		a[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if a[i] == n-1 {
			return i
		}
	}
	return -1
}
