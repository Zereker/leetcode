package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var result int
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		result = f1 + f2
		f1 = f2
		f2 = result
	}

	return result
}

func main() {
	println(climbStairs(3))
}
