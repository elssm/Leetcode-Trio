package goLeetCode

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lowx, highx := 0, m-1
	lowy, highy := 0, n-1
	for lowx <= highx && lowy <= highy {
		midx := lowx + (highx-lowx)/2
		midy := lowy + (highy-lowy)/2
		if matrix[midx][midy] == target {
			return true
		}
		if matrix[midx][midy] > target {
			if matrix[midx][0] == target {
				return true
			}
			if matrix[midx][0] > target {
				highx = midx - 1
			} else {
				highy = midy - 1
			}
		} else {
			if matrix[midx][n-1] == target {
				return true
			}
			if matrix[midx][n-1] > target {
				lowy = midy + 1
			} else {
				lowx = midx + 1
			}
		}
	}
	return false
}
