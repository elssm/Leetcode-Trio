package goLeetCode

func wordBreak2(s string, wordDict []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	var res []string
	var tmp []string
	var f func(s string, tmp []string)
	f = func(s string, tmp []string) {
		if s == "" {
			str := ""
			for _, v := range tmp {
				str += v
				str += " "
			}
			str = str[:len(str)-1]
			res = append(res, str)
			return
		}
		for i := 1; i <= len(s); i++ {
			if wordMap[s[:i]] {
				tmp = append(tmp, s[:i])
				f(s[i:], tmp)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	f(s, tmp)
	return res
}
