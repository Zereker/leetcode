package main

import "math"

func myAtoi(str string) int {
	positive, start := isPositive(str)

	var chars []rune
	runes := []rune(str)
	for i := start; i < len(runes); i++ {
		if str[i] >= '0' && runes[i] <= '9' {
			chars = append(chars, runes[i])
		} else {
			break
		}
	}

	return runesTo(positive, chars)
}

func runesTo(positive bool, runes []rune) int {
	var sum int

	for i := 0; i < len(runes); i++ {
		digit := runes[i] - '0'
		if math.MaxInt32/10 < sum || math.MaxInt32/10 == sum && math.MaxInt32%10 < digit {
			if positive {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		sum = 10*sum + int(digit)
	}

	if positive {
		return sum
	} else {
		return sum * -1
	}
}

func isPositive(str string) (bool, int) {
	for index, char := range str {
		if char == ' ' {
			continue
		} else if char == '+' || char == '-' {
			return char == '+', index + 1
		} else {
			return true, index
		}
	}

	return true, 0
}

func main() {
	println(myAtoi("42"))
	println(myAtoi("   -42"))
	println(myAtoi("4193 with words"))
	println(myAtoi("words and 987"))
	println(myAtoi("-91283472332"))
}
