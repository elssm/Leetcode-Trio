package goLeetCode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var res = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
