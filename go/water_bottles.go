package main

func numWaterBottles(numBottles int, numExchange int) int {
	r := numBottles
	e := numBottles
	for {
		if f := e / numExchange; f == 0 {
			return r
		} else {
			e = e%numExchange + f
			r += f
		}
	}
}
