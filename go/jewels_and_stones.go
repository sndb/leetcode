package main

func numJewelsInStones(jewels string, stones string) int {
	j := map[rune]bool{}
	for _, v := range jewels {
		j[v] = true
	}
	c := 0
	for _, v := range stones {
		if j[v] {
			c++
		}
	}
	return c
}
