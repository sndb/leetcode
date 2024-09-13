package main

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	for i, n := range arr {
		prefix[i+1] = prefix[i] ^ n
	}

	result := make([]int, 0, len(queries))
	for _, q := range queries {
		result = append(result, prefix[q[0]]^prefix[q[1]+1])
	}
	return result
}
