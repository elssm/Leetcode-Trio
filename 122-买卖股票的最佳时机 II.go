package goLeetCode

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyP, res := prices[0], 0
	for _, v := range prices {
		if v < buyP {
			buyP = v
		} else if v > buyP {
			res += v - buyP
			buyP = v
		}
	}
	return res
}
