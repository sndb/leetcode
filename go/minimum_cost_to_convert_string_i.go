package main

import (
	"fmt"
	"math"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	g := make([][]int, 26)
	for i := range g {
		g[i] = make([]int, len(g))
		for j := range g[i] {
			if i != j {
				g[i][j] = math.MaxInt
			}
		}
	}
	for k, c := range cost {
		i := original[k] - 'a'
		j := changed[k] - 'a'
		g[i][j] = min(g[i][j], c)
	}
	for k := range g {
		for i := range g {
			for j := range g {
				if g[i][j]-g[i][k] > g[k][j] {
					g[i][j] = g[i][k] + g[k][j]
				}
			}
		}
	}
	var total int64 = 0
	for k := range source {
		c := g[source[k]-'a'][target[k]-'a']
		if c == math.MaxInt {
			return -1
		}
		total += int64(c)
	}
	return total
}

func main() {
	fmt.Println(minimumCost("abcd", "acbe", []byte{'a', 'b', 'c', 'c', 'e', 'd'}, []byte{'b', 'c', 'b', 'e', 'b', 'e'}, []int{2, 5, 5, 1, 2, 20}))
}
