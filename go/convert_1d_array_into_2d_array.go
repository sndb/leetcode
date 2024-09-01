package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	result := make([][]int, 0, m)
	for i := range m {
		result = append(result, original[n*i:n*i+n])
	}
	return result
}
