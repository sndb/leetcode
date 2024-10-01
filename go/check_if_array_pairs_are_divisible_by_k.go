package main

func canArrange(arr []int, k int) bool {
	r := 0
	h := map[int]int{}
	for _, v := range arr {
		x := (v%k + k) % k
		y := (k - x) % k
		if h[y] > 0 {
			h[y]--
			r++
		} else {
			h[x]++
		}
	}
	return r == len(arr)/2
}
