package main

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 0)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {

}
