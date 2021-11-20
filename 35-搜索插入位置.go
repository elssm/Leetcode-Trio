package goLeetCode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		numMap1 := make(map[byte]int)
		numMap2 := make(map[byte]int)
		for j := 0; j < len(board); j++ {
			rv := board[i][j]
			cv := board[j][i]
			if i%3 == 0 && j%3 == 0 {
				numMap3 := make(map[byte]int)
				for ii := i; ii < i+3; ii++ {
					for jj := j; jj < j+3; jj++ {
						sanv := board[ii][jj]
						if sanv != '.' {
							if _, ok := numMap3[sanv]; ok {
								return false
							} else {
								numMap3[sanv] = 1
							}
						}
					}
				}
			}
			if rv != '.' {
				if _, ok := numMap1[rv]; ok {
					return false
				} else {
					numMap1[rv] = 1
				}
			}
			if cv != '.' {
				// fmt.Printf("%c ", cv)
				if _, ok := numMap2[cv]; ok {
					return false
				} else {
					numMap2[cv] = 1
				}
			}
		}
	}
	return true
}
