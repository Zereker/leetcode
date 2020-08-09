package main

func fib(N int) int {
	dp := []int{0, 1}

	if N <= 1 {
		return dp[N]
	}

	for i := 2; i <= N; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[N]
}

func main() {
	println(fib(100))
}
