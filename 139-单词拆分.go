package goLeetCode

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			if wordMap[suffix] && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
