package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if 1<<maxChoosableInteger <= desiredTotal {
		return false
	}
	var f, g func(int, int) byte
	f = func(t, b int) byte {
		for i := 0; i < maxChoosableInteger; i++ {
			nb := b | 1<<i
			nt := t + i + 1
			if nb != b && (nt >= desiredTotal || g(nt, nb) == 2) {
				return 1
			}
		}
		return 2
	}
	m := make([]byte, 1<<maxChoosableInteger)
	g = func(t, b int) byte {
		if m[b] == 0 {
			m[b] = f(t, b)
		}
		return m[b]
	}
	return f(0, 0) == 1
}

func main() {
	fmt.Println(canIWin(10, 40)) // false
	fmt.Println(canIWin(4, 6))   // true
	fmt.Println(canIWin(7, 16))  // true
}
