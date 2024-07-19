package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if i-1 >= 0 && flowerbed[i-1] == 1 {
			continue
		}
		if flowerbed[i] == 1 {
			continue
		}
		if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
			continue
		}
		flowerbed[i] = 1
		n--
	}
	return n == 0
}
