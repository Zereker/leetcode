package main

import "strings"

func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")

	length := len(split)
	if length == 0 {
		return len(s)
	}

	for i := length - 1; i >= 0; i-- {
		if len(split[i]) != 0 {
			return len(split[i])
		}
	}

	return 0
}

func main() {
	println(lengthOfLastWord("a a a a a a a a       "))
}
