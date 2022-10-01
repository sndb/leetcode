package main

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	buckets := make([][]int, len(nums)+1)
	for x, c := range freq {
		buckets[c] = append(buckets[c], x)
	}

	r := make([]int, 0, k)
outer:
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, x := range buckets[i] {
			if len(r) == cap(r) {
				break outer
			}
			r = append(r, x)
		}
	}
	return r
}
