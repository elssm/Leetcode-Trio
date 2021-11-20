package goLeetCode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	i1, i2 := 1, 2
	for i := 3; i <= n; i++ {
		temp := i1 + i2
		i1, i2 = i2, temp
	}
	return i2
}

func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[0] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
