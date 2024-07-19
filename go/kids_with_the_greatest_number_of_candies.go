package main

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := slices.Max(candies)
	result := make([]bool, len(candies))
	for i, c := range candies {
		result[i] = c+extraCandies >= max
	}
	return result
}
