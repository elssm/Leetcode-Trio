package goLeetCode

func numDecodings(s string) int {
	if len(s) == 0 || (len(s) == 1 && s[0] == '0') {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			dp[i+1] = 0
		} else {
			dp[i+1] = dp[i]
		}
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && int(s[i]-'0') <= 6)) {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}
