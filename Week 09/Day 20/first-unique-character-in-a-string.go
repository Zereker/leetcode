package main

func firstUniqChar(s string) int {
	var chars [26]int

	for index, char := range s {
		chars[char-'a'] = index
	}

	for index, char := range s {
		if chars[char-'a'] == index {
			return index
		} else {
			chars[char-'a'] = -1
		}
	}

	return -1
}

func main() {
	uniqChar := firstUniqChar("loveleetcode")
	println(uniqChar)
}
