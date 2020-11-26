package main

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return 1
	}

	var result []int
	for i := 0; i < len(ratings); i++ {
		result = append(result, 1)
	}

	for i := 0; i < len(ratings)-1; i++ {
		left := ratings[i]
		right := ratings[i+1]

		if left < right {
			result[i+1] = result[i] + 1
		}
	}

	for i := len(ratings) - 1; i > 0; i-- {
		left := ratings[i-1]
		right := ratings[i]

		// 注意: 第二次遍历的时候, 有可能左边增加后的数量仍然小于右边增加后的数量
		if left > right && result[i-1] <= result[i] {
			result[i-1] = result[i] + 1
		}
	}

	var sum int
	for _, i := range result {
		sum += i
	}

	return sum
}

func main() {
	println(candy([]int{1, 0, 2}))
	println(candy([]int{1, 2, 2}))
	println(candy([]int{1, 3, 4, 5, 2})) // 1, 2, 3, 4, 1
}
