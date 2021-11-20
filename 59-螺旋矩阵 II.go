package goLeetCode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	i, j, di, dj := 0, 0, 0, 1
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for v := 1; v <= n*n; v++ {
		matrix[i][j] = v
		if matrix[(i+di+n)%n][(j+dj+n)%n] != 0 {
			di, dj = dj, -di
		}
		i += di
		j += dj
	}
	return matrix
}
