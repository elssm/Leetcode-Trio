package goLeetCode

func isScramble(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if len(s1) != len(s2) {
		return false
	}
	dp := make([][][]bool, m)
	for i := range dp {
		dp[i] = make([][]bool, m)
		for j := range dp[i] {
			dp[i][j] = make([]bool, m+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for k := 1; k <= l-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][m]
}
