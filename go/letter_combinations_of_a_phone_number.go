package main

import "fmt"

var letters = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return letters[digits[0]]
	}
	var r []string
	for _, c := range letterCombinations(digits[1:]) {
		for _, l := range letters[digits[0]] {
			r = append(r, l+c)
		}
	}
	return r
}

func main() {
	fmt.Println(letterCombinations("23"))
}
