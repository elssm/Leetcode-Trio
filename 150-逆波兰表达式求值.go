package goLeetCode

func evalRPN(tokens []string) int {
	var numStack []int
	for _, v := range tokens {
		// fmt.Println(numStack)
		if v != "+" && v != "-" && v != "*" && v != "/" {
			numStack = append(numStack, toNumber(v))
		} else {
			a := numStack[len(numStack)-1]
			b := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-2]
			switch v {
			case "+":
				numStack = append(numStack, a+b)
			case "-":
				numStack = append(numStack, b-a)
			case "*":
				numStack = append(numStack, a*b)
			case "/":
				numStack = append(numStack, b/a)
			default:
			}
		}
	}
	return numStack[len(numStack)-1]
}

func toNumber(s string) int {
	flag := 1
	if s[0] == '-' {
		flag = -1
		s = s[1:]
	}
	ans := 0
	for _, v := range s {
		ans = int(v-'0') + ans*10
	}
	return ans * flag
}
