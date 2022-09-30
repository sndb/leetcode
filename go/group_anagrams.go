package main

import "sort"

func sortString(s string) string {
	r := []byte(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		key := sortString(v)
		m[key] = append(m[key], v)
	}
	r := make([][]string, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
