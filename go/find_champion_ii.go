package main

func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true
	}
	champ := -1
	for i := range n {
		if weak[i] {
			continue
		}
		if champ != -1 {
			return -1
		}
		champ = i
	}
	return champ
}
