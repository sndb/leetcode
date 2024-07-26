package main

import (
	"fmt"
	"math"
)

type road struct {
	dest   int
	weight int
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	g := map[int][]road{}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], road{v, w})
		g[v] = append(g[v], road{u, w})
	}
	xc, xr := -1, math.MaxInt
	for c := 0; c < n; c++ {
		if r := countReachable(g, c, distanceThreshold); r <= xr {
			xc, xr = c, r
		}
	}
	return xc
}

func countReachable(g map[int][]road, o, t int) int {
	best := map[int]int{}
	var walk func(int, int)
	walk = func(u int, w int) {
		for _, v := range g[u] {
			if x := w + v.weight; v.dest != o && x <= t {
				if b, ok := best[v.dest]; !ok || x < b {
					best[v.dest] = x
					walk(v.dest, x)
				}
			}
		}
	}
	walk(o, 0)
	return len(best)
}

func main() {
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
}
