package goLeetCode

func numDistinct1(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if j > i {
				continue
			}
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

func numDistinct2(s string, t string) int {
	m, n := len(s), len(t)
	if n > m {
		return 0
	}
	if n == 0 {
		return 1
	}
	// dp[i][j] 是 s[:i-1] 中 找到子序列等于 t[:j-1] 的个数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			// 这两个base顺序不能变更
			if j == 0 {
				// t 变成空串，此时 s 为了匹配它，方式只有1种：一个字符都不挑。
				dp[i][j] = 1
			} else if i == 0 {
				// s 变成空串，但 t 不是，s 怎么也匹配不了 t，方式数为 0
				dp[i][j] = 0
			} else {
				if s[i-1] == t[j-1] {
					dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}
