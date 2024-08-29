package main

func removeStones(stones [][]int) int {
	ys := map[int][]int{}
	xs := map[int][]int{}
	for _, s := range stones {
		x, y := s[0], s[1]
		ys[x] = append(ys[x], y)
		xs[y] = append(xs[y], x)
	}

	seen := map[[2]int]bool{}
	var remove func(int, int) int
	remove = func(x, y int) int {
		k := [2]int{x, y}
		if seen[k] {
			return 0
		}
		seen[k] = true

		n := 0
		for _, nx := range xs[y] {
			n += remove(nx, y)
		}
		for _, ny := range ys[x] {
			n += remove(x, ny)
		}
		return n + 1
	}

	r := 0
	for _, s := range stones {
		r += max(0, remove(s[0], s[1])-1)
	}
	return r
}
