package goLeetCode

import (
	"fmt"
	"strconv"
)

func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	res := ""
	str := countAndSay1(n - 1)
	diff := str[0]
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == diff {
			count++
		} else {
			res += strconv.Itoa(count) + string(diff)
			count = 1
			diff = str[i]
		}
	}
	res += strconv.Itoa(count) + string(diff)
	return res
}

func countAndSay2(n int) string {
	res := make([]byte, 1)
	res[0] = '1'
	if n == 1 {
		return string(res)
	}
	var a byte
	var c byte
	count := 0
	tem := make([]byte, 0)
	for i := 2; i <= n; i++ {
		a = res[0]
		count = 0
		for _, c = range res {
			if c == a {
				count++
			} else {
				tem = append(tem, byte(count+48), a)
				a = c
				count = 1
			}
		}
		tem = append(tem, byte(count+48), a)
		res = tem
		tem = make([]byte, 0)
		fmt.Println(string(res))
	}
	return string(res)
}
