package goLeetCode

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = dp[i+1][n-1] + grid[i][n-1]
	}
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = dp[m-1][i+1] + grid[m-1][i]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
