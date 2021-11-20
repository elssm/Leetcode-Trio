package goLeetCode

var tar float64

func f(x float64) int {
	res := x*x - tar
	if res < 1e-1 {
		return int(x)
	}
	x = x - res/(2*x)
	return f(x)

}
func mySqrt(x int) int {
	tar = float64(x)

	return f(float64(x))
}
