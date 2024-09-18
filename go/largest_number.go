package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	r := []string{}
	for _, n := range nums {
		r = append(r, strconv.Itoa(n))
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i]+r[j] > r[j]+r[i]
	})
	if r[0][0] == '0' {
		return "0"
	}
	return strings.Join(r, "")
}
