package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		n--
		return n <= 0
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i == 0 { //左边界
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else if i == len(flowerbed)-1 { //右边界
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else {
				if (flowerbed[i+1] == 0) && (flowerbed[i-1] == 0) {
					flowerbed[i] = 1
					n--
				}
			}
		}
	}

	return n <= 0
}

func main() {
	println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
