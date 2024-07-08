package main

func findTheWinner(n int, k int) int {
	l := make([]int, n)
	for i := range n {
		l[i] = i + 1
	}
	i := 0
	for len(l) > 1 {
		i = (i + k - 1) % len(l)
		l = append(l[:i], l[i+1:]...)
	}
	return l[0]
}
