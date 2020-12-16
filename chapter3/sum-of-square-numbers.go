package main

import "math"

func judgeSquareSum(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))

	for a <= b {
		sum := (a * a) + (b * b)

		if sum < c {
			a++
		} else if sum > c {
			b--
		} else {
			return true
		}
	}

	return false
}

func main() {
	println(judgeSquareSum(5))
	println(judgeSquareSum(3))
	println(judgeSquareSum(4))
	println(judgeSquareSum(2))
	println(judgeSquareSum(999999999))
}
