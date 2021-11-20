package goLeetCode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	reach, nextReach := 0, nums[0]
	res := 0
	for i := 0; i < len(nums); i++ {
		nextReach = max(nums[i]+i, nextReach)
		if nextReach >= len(nums)-1 {
			return res + 1
		}
		if i == reach {
			res++
			reach = nextReach
		}
	}
	return res
}
