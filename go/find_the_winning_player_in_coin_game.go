package main

func losingPlayer(x int, y int) string {
	i := 1
	for x >= 1 && y >= 4 {
		x -= 1
		y -= 4
		i ^= 1
	}
	return [...]string{"Alice", "Bob"}[i]
}
