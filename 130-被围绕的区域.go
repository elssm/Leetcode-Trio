package goLeetCode

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if board[i][j] == 'O' {
					dfsSolve(i, j, board)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsSolve(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		dfsSolve(i-1, j, board)
		dfsSolve(i+1, j, board)
		dfsSolve(i, j-1, board)
		dfsSolve(i, j+1, board)
	}
}
