package main

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] == s[right] {
			left++
			right--
		} else {

		}

	}

	return false
}

func main() {

}
