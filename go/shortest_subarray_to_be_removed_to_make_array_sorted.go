package main

func findLengthOfShortestSubarray(arr []int) int {
	l := 0
	r := len(arr) - 1
	for r > 0 && arr[r-1] <= arr[r] {
		r--
	}

	x := r
	for l < r && (l == 0 || arr[l-1] <= arr[l]) {
		for r < len(arr) && arr[l] > arr[r] {
			r++
		}
		x = min(x, r-l-1)
		l++
	}
	return x
}
