package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	mSum := 0
	for _, x := range rolls {
		mSum += x
	}
	nSum := mean*(n+len(rolls)) - mSum

	result := make([]int, 0, n)
	for n > 0 {
		x := nSum / n
		if x < 1 || x > 6 {
			return nil
		}
		result = append(result, x)
		nSum -= x
		n--
	}
	return result
}

func main() {
	fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))
	fmt.Println(missingRolls([]int{1, 2, 3, 4}, 6, 4))
}
