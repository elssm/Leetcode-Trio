package goLeetCode

import "fmt"

func singleNumber2(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%x, %x", ones, twos)
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
