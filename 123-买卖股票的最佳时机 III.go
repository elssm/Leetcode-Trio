package goLeetCode

import (
	"fmt"
	"math"
)

func maxProfit2(prices []int) int {
	fstBuy, fstSell := math.MinInt32, 0
	secBuy, secSell := math.MinInt32, 0
	for _, v := range prices {
		fstBuy = max(fstBuy, -v)
		fstSell = max(fstSell, fstBuy+v)
		secBuy = max(secBuy, fstSell-v)
		secSell = max(secSell, secBuy+v)
		fmt.Println(fstBuy, fstSell, secBuy, secSell)
	}
	return secSell
}


