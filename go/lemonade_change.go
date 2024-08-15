package main

func lemonadeChange(bills []int) bool {
	n5, n10 := 0, 0
	for _, b := range bills {
		if b == 5 {
			n5++
		} else if b == 10 {
			if n5 > 0 {
				n5--
				n10++
			} else {
				return false
			}
		} else {
			if n10 > 0 && n5 > 0 {
				n5--
				n10--
			} else if n5 > 2 {
				n5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
