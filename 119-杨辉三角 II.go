package goLeetCode

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for j := 1; j <= rowIndex; j++ {
		for i := j; i >= 0; i-- {
			if i >= 1 {
				res[i] = res[i-1] + res[i]
			}
		}
	}
	return res
}
