package goLeetCode

func generate(numRows int) [][]int {
	var ans [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, row)
	}
	return ans
}
