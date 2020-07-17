package main

import (
	"github.com/Zereker/LeetCode/containers"
)

var parentheses = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := containers.NewStack()

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.Push(r)
		case ')', '}', ']':
			pop, err := stack.Pop()
			if err != nil {
				return false
			}

			r2 := pop.(rune)
			if r2 != parentheses[r] {
				return false
			}
		case ' ':
			continue
		default:
			return false
		}
	}

	if !stack.Empty() {
		return false
	}

	return true
}

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
}
