package goLeetCode

func exist(board [][]byte, word string) bool {
	mark := make([][]bool, len(board))
	for i := range mark {
		mark[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if isExist(&board, mark, i, j, word) {
				return true
			}
		}
	}
	return false
}

func isExist(board *[][]byte, mark [][]bool, i, j int, word string) bool {
	if word == "" {
		return true
	}
	if (*board)[i][j] == word[0] && mark[i][j] == false {
		// fmt.Println(word)
		le, ri, up, down := false, false, false, false
		mark[i][j] = true
		word = word[1:]
		if i-1 >= 0 {
			le = isExist(board, mark, i-1, j, word)
		}
		if i+1 < len(*board) {
			ri = isExist(board, mark, i+1, j, word)
		}
		if j-1 >= 0 {
			up = isExist(board, mark, i, j-1, word)
		}
		if j+1 < len((*board)[0]) {
			down = isExist(board, mark, i, j+1, word)
		}
		mark[i][j] = false
		return le || ri || up || down || word == ""
	}
	return false
}
