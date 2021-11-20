package goLeetCode

func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			Mswap(&matrix, i, j, n)
		}
	}
}
func Mswap(matrix *[][]int, i, j, n int) {
	(*matrix)[i][j], (*matrix)[j][n-1-i], (*matrix)[n-1-i][n-1-j], (*matrix)[n-1-j][i] = (*matrix)[n-1-j][i], (*matrix)[i][j], (*matrix)[j][n-1-i], (*matrix)[n-1-i][n-1-j]
}

func rotate2(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
