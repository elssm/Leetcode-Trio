package goLeetCode

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		middle := (i + j) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			// 三种情况
			if nums[middle] > nums[i] && nums[middle] > nums[j] {
				i = middle + 1
			} else if nums[middle] <= nums[i] && nums[middle] <= nums[j] {
				if target > nums[j] {
					j = middle - 1
				} else if target < nums[j] {
					i = middle + 1
				} else {
					return j
				}
			} else {
				i = middle + 1
			}
		} else {
			// 三种情况
			if nums[middle] >= nums[i] && nums[middle] >= nums[j] {
				if target > nums[i] {
					j = middle - 1
				} else if target < nums[i] {
					i = middle + 1
				} else {
					return i
				}
			} else if nums[middle] < nums[i] && nums[middle] < nums[j] {
				j = middle - 1
			} else {
				j = middle - 1
			}
		}
	}
	return -1
}
