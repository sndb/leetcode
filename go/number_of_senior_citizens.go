package main

import "strconv"

func countSeniors(details []string) int {
	r := 0
	for _, d := range details {
		n, _ := strconv.Atoi(d[11:13])
		if n > 60 {
			r++
		}
	}
	return r
}
