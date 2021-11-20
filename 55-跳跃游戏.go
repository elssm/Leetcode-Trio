package goLeetCode

func canJump(nums []int) bool {
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}
		if i == 0 && n > 1 {
			return false
		}
	}
	return true
}
