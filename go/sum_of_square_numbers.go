package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))
	for i <= j {
		n := i*i + j*j
		if n < c {
			i++
		} else if n > c {
			j--
		} else {
			return true
		}
	}
	return false
}

func main() {
	for i := 0; i < 16; i++ {
		fmt.Println(i, judgeSquareSum(i))
	}
}
