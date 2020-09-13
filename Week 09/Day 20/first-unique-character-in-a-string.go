package main

func firstUniqChar(s string) int {
	m := make(map[rune]int, len(s))
	for _, i := range s {
		if m[i] == 0 {
			m[i] += 1
		}
	}

	var char rune
	for r, i := range m {
		if i == 1 {
			char = r
			break
		}
	}


	for i, i2 := range s {
		if i2 == char {
			return i
		}
	}

	return 0
}

func main() {
	uniqChar := firstUniqChar("leetcode")
	println(uniqChar)
}
