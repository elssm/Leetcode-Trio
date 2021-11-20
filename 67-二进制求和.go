package goLeetCode

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m < n {
		a, b = b, a
		m, n = n, m
	}
	res := ""
	jinw := 0
	for i, j := m-1, n-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			res = string((int(a[i]-'0')+int(b[j]-'0')+jinw)%2+'0') + res
			jinw = (int(a[i]-'0') + int(b[j]-'0') + jinw) / 2
		} else {
			res = string((int(a[i]-'0')+jinw)%2+'0') + res
			jinw = (int(a[i]-'0') + jinw) / 2

		}
	}
	if jinw > 0 {
		res = "1" + res
	}
	return res
}
