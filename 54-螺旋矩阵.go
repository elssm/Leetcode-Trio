package goLeetCode

func spiralOrder(matrix [][]int) []int {
	r, i, j, di, dj := []int{}, 0, 0, 0, 1
	m, n := len(matrix), len(matrix[0])
	for k := 0; k < len(matrix)*len(matrix[0]); k++ {
		r = append(r, matrix[i][j])
		matrix[i][j] = 0
		if matrix[(i+di+m)%m][(j+dj+n)%n] == 0 {
			di, dj = dj, -di
		}
		i += di
		j += dj
	}
	return r
}
