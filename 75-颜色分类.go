package goLeetCode

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
	return
}

func partition(nums []int, low, hight int) int {
	pivot := nums[hight]
	mark := low - 1
	for i := low; i < hight; i++ {
		if nums[i] < pivot {
			mark++
			nums[mark], nums[i] = nums[i], nums[mark]
		}
	}
	if nums[hight] < nums[mark+1] {
		nums[mark+1], nums[hight] = nums[hight], nums[mark+1]
	}
	return mark + 1
}

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, lo, hi)
		quickSort(nums, lo, p-1)
		quickSort(nums, p+1, hi)
	}

}
