package main

import (
	"fmt"
	"math"
	"strconv"
)

func nearestPalindromic(s string) string {
	n := len(s)
	whole, _ := strconv.ParseInt(s, 10, 64)
	half, _ := strconv.ParseInt(s[:n/2+n%2], 10, 64)
	odd := n%2 != 0
	cands := []int64{
		mirror(half+1, odd),
		mirror(half, odd),
		mirror(half-1, odd),
		zero(whole) - 1,
		zero(whole)*10 + 1,
	}
	best := int64(math.MaxInt64)
	diff := int64(math.MaxInt64)
	for _, cand := range cands {
		if d := abs(cand, whole); d != 0 && d <= diff {
			best = cand
			diff = d
		}
	}
	return strconv.FormatInt(best, 10)
}

func abs(x, y int64) int64 {
	d := x - y
	if d < 0 {
		d = -d
	}
	return d
}

func mirror(x int64, odd bool) int64 {
	y := int64(0)
	z := x
	if odd {
		x /= 10
	}
	for z > 0 {
		y *= 10
		y += z % 10
		x *= 10
		z /= 10
	}
	return x + y
}

func zero(x int64) int64 {
	y := int64(1)
	for x > 9 {
		x /= 10
		y *= 10
	}
	return y
}

func main() {
	fmt.Println(nearestPalindromic("12345"))
	fmt.Println(nearestPalindromic("10"))
	fmt.Println(nearestPalindromic("11"))
}
