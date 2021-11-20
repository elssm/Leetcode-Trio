package goLeetCode

import "strings"

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	pos := 0
	res := make([]byte, l1+l2+1) // 可能向前进位，多保留一位
	for i := range res {         // 初始化res
		res[i] = '0'
	}
	for i := l1 - 1; i >= 0; i-- { // 需要用一个数乘以另一个数的每一位，因此两层遍历
		for j := l2 - 1; j >= 0; j-- {
			temp := (num1[i] - '0') * (num2[j] - '0')
			pos = i + j + 2
			res[pos-1] += (temp + res[pos] - '0') / 10 // 进位
			res[pos] = (temp+res[pos]-'0')%10 + '0'
		}
	}
	ret := strings.TrimLeft(string(res), "0") // 移除左侧多余的0
	if ret == "" {
		return "0"
	}
	return ret
}
