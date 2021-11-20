package goLeetCode

import "strings"

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	res := ""
	var ans []string
	for _, s := range strs {
		if s == "." || s == "" {
			continue
		}
		if s == ".." {
			if len(ans) != 0 {
				ans = ans[:len(ans)-1]
			}
			continue
		}
		ans = append(ans, s)
	}
	if len(ans) == 0 {
		return "/"
	}
	for _, v := range ans {
		res += "/" + v
	}
	return res
}
