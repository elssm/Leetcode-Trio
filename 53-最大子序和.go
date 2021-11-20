package goLeetCode

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if res < sum {
			res = sum
		}
	}
	return res
}
