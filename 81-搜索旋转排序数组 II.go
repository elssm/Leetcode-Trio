package goLeetCode

func search1(nums []int, target int) bool {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] == target || nums[i] == target || nums[j] == target {
			return true
		}
		if nums[i] == nums[j] {
			i++
			j--
		} else {
			if nums[i] <= nums[m] {
				if nums[i] < target && target < nums[m] {
					j = m - 1
				} else {
					i = m + 1
				}
			} else {
				if nums[m] < target && target < nums[j] {
					i = m + 1
				} else {
					j = m - 1
				}
			}
		}
	}
	return false
}
