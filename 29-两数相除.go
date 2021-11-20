package goLeetCode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	res := 0
	sign := !((divisor > 0) != (dividend > 0))
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			res += 1 << i
			dividend -= divisor << i
		}
	}
	if sign {
		return res
	}
	return -res
}
