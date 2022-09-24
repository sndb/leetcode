package main

func climbStairsCache(n int, cache map[int]int) int {
	x, ok := cache[n]
	if ok {
		return x
	}
	cache[n-1] = climbStairsCache(n-1, cache)
	cache[n-2] = climbStairsCache(n-2, cache)
	return cache[n-2] + cache[n-1]
}

func climbStairs(n int) int {
	cache := map[int]int{1: 1, 2: 2}
	return climbStairsCache(n, cache)
}
