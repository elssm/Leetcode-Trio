package goLeetCode

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if counter[uniqNums[i]] >= 4 && uniqNums[i]*4 == target {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if counter[uniqNums[i]] >= 3 && uniqNums[i]*3+uniqNums[j] == target {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 && uniqNums[i]*2+uniqNums[j]*2 == target {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			if counter[uniqNums[j]] >= 3 && uniqNums[j]*3+uniqNums[i] == target {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if counter[uniqNums[i]] >= 2 && uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if counter[uniqNums[j]] >= 2 && uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if counter[uniqNums[k]] >= 2 && uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				for l := k + 1; l < len(uniqNums); l++ {
					if uniqNums[i]+uniqNums[j]+uniqNums[k]+uniqNums[l] == target {
						res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[l]})
					}
				}
			}
		}
	}
	return res
}
