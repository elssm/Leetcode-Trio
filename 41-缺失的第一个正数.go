package goLeetCode

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] > len(nums) || nums[i] == nums[nums[i]-1] {
				break
			}
			// 将nums[i] 放置到对应位置上[1,2,3...]
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}
	return len(nums) + 1

}
