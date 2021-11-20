package goLeetCode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var l, r = -1, -1
	var res []int
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				l = mid
				break
			}
			j = mid - 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	// fmt.Println(l)
	if l == -1 {
		return []int{-1, -1}
	}
	for i := l; i < len(nums); i++ {
		if nums[i] == target {
			r = i
		} else {
			break
		}
	}
	res = append(res, l, r)
	return res

}
