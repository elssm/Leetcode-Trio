package goLeetCode

func removeDuplicates1(nums []int) int {
	index := 0
	for _, v := range nums {
		if index < 2 || v > nums[index-2] {
			nums[index] = v
			index++
		}
	}
	return index
}
