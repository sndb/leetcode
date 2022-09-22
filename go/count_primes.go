package main

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primes := make([]bool, n)
	x := 0
	for i := 2; i < n; i++ {
		if !primes[i] {
			for j, v := i, 1; j*v < n; v++ {
				primes[j*v] = true
			}
			x++
		}
	}
	return x
}

func main() {
	fmt.Println(countPrimes(10))
}
