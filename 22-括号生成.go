package goLeetCode

func generateParenthesis(n int) []string {
	var res []string
	generatedValidBracket(n, n, "", &res)
	return res
}

func generatedValidBracket(LIndex int, RIndex int, s string, res *[]string){
	if LIndex == 0 && RIndex == 0{
		*res = append(*res, s)
		return
	}
	if LIndex > 0 {
		generatedValidBracket(LIndex- 1, RIndex, s + "(", res)
	}
	if LIndex < RIndex {
		generatedValidBracket(LIndex, RIndex-1, s + ")", res)
	}
}
