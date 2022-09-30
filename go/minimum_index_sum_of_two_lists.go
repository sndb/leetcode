package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	s1 := map[string]int{}
	for i, s := range list1 {
		s1[s] = i
	}
	r := map[int][]string{}
	min := math.MaxInt
	for i, s := range list2 {
		j, ok := s1[s]
		if !ok {
			continue
		}
		v := i + j
		if v <= min {
			min = v
			r[v] = append(r[v], s)
		}
	}
	return r[min]
}
