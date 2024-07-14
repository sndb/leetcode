package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	k := 0
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		count := 1
		for i+1 < len(chars) && chars[i+1] == char {
			count++
			i++
		}
		chars[k] = char
		k++
		if count > 1 {
			digits := strconv.Itoa(count)
			for j := 0; j < len(digits); j++ {
				chars[k] = digits[j]
				k++
			}
		}
	}
	return k
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	fmt.Println(string(chars))
}
