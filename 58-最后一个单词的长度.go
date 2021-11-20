package goLeetCode

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s) - 1 - i
		}
	}
	return len(s)
}
