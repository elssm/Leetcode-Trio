package goLeetCode

import "math"

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	for i := range s {
		if dp[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 1; j <= i; j++ {
			if dp[j][i] && f[j-1]+1 < f[i] {
				f[i] = f[j-1] + 1
			}
		}
	}
	return f[n-1]
}
