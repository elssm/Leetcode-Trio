package goLeetCode

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	addr, res := "", []string{}
	findIpAddresses(s, addr, &res, 0)
	return res
}

func findIpAddresses(s, addr string, res *[]string, index int) {
	if index == 3 {
		vl, _ := strconv.Atoi(s)
		if len(s) == 1 || (s[0] != '0' && vl <= 255) {
			addr += s
			*res = append(*res, addr)
		}
		return
	}
	for i := 1; i < len(s); i++ {
		tmp := s[:i]
		if len(tmp) == 1 || tmp[0] != '0' {
			if val, _ := strconv.Atoi(tmp); val <= 255 {
				addr += s[:i] + "."
				findIpAddresses(s[i:], addr, res, index+1)
				addr = addr[:len(addr)-(i+1)]
			}
		}

	}
}
