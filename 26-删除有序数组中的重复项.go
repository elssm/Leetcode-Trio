package goLeetCode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var comp = nums[0]
	var res = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != comp {
			res++
			nums[res] = nums[i]
			comp = nums[i]
		}
	}
	return res + 1
}
