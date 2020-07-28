package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	recursive(0, 0, n, "", &result)
	return result
}

func recursive(left, right int, n int, s string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, s)
	}

	if left < n {
		recursive(left+1, right, n, s+"(", result)
	}

	if right < left {
		recursive(left, right+1, n, s+")", result)
	}
}

func main() {
	fmt.Printf("%v", generateParenthesis(3))
}
