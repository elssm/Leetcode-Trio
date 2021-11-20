package goLeetCode

func maxProfit(prices []int) int {
	maxp := 0
	minp := 1<<63 - 1
	for i := range prices {
		if prices[i] < minp {
			minp = prices[i]
		} else if prices[i]-minp > maxp {
			maxp = prices[i] - minp
		}
	}
	return maxp
}
