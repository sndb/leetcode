package main

import "fmt"

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	ret := [][]int{{rStart, cStart}}
	dir := 0
	run := 0
	for len(ret) < rows*cols {
		for range run/2 + 1 {
			switch dir {
			case 0:
				cStart++
			case 1:
				rStart++
			case 2:
				cStart--
			case 3:
				rStart--
			}
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				ret = append(ret, []int{rStart, cStart})
			}
		}
		dir = (dir + 1) % 4
		run++
	}
	return ret
}

func main() {
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}
