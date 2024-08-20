package main

import (
	"fmt"
	"math"
)

func stoneGameII(piles []int) int {
	cache := make([][]int, len(piles))
	for i := range cache {
		cache[i] = make([]int, len(piles))
	}
	var f func(int, int) int
	f = func(p, m int) int {
		stones := 0
		if p+2*m >= len(piles) {
			for i := p; i < len(piles); i++ {
				stones += piles[i]
			}
			return stones
		}
		if v := cache[p][m]; v != 0 {
			return v
		}
		best := 0
		for i := 1; i <= 2*m; i++ {
			stones += piles[p+i-1]
			worst := math.MaxInt
			for j := 1; j <= 2*max(m, i); j++ {
				worst = min(worst, f(p+i+j, max(m, i, j)))
			}
			best = max(best, stones+worst)
		}
		cache[p][m] = best
		return best
	}
	return f(0, 1)
}

func main() {
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
	fmt.Println(stoneGameII([]int{1, 2, 3, 4, 5, 100}))
	fmt.Println(stoneGameII([]int{7468, 6245, 9261, 3958, 1986, 1074, 5677, 9386, 1408, 1384, 8811, 3885, 9678, 8470, 8893, 7514, 4941, 2148, 5217, 5425, 5307, 747, 1253, 3518, 5238, 5834, 9133, 8391, 6100, 3362, 7807, 2581, 6121, 7684, 8744, 9584, 4068, 7204, 4285, 8635}))
	fmt.Println(stoneGameII([]int{1}))
}
