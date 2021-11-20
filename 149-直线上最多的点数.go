package goLeetCode

func maxPoints(points [][]int) int {
	res := 1
	if len(points) == 1 {
		return res
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			star := points[i]
			incx := points[i][0] - points[j][0]
			incy := points[i][1] - points[j][1]
			// fmt.Println(incx, incy)
			tmp := 0
			for _, p := range points {
				if (p[0]-star[0])*incy == (p[1]-star[1])*incx {
					tmp++
				}
			}
			res = max(res, tmp)
		}
	}
	return res
}

