package goLeetCode

func partition2(s string) [][]string {
	res, tmp := [][]string{}, []string{}
	findPartition(s, tmp, &res)
	return res
}

func findPartition(s string, tmp []string, res *[][]string) {
	if s == "" {
		tt := make([]string, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	}
	for i := 1; i <= len(s); i++ {
		if isPalin(s[:i]) {
			tmp = append(tmp, s[:i])
			findPartition(s[i:], tmp, res)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func isPalin(s string) bool {
	return s == Reverse(s)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
