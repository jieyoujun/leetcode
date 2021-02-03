package validsudoku

func IsValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	return row(board) && col(board) && cell(board)
}

func row(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := [9]int{}
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				if m[board[i][j]-'1'] == 0 {
					m[board[i][j]-'1']++
				} else {
					return false
				}
			}
		}
	}
	return true
}

func col(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := [9]int{}
		for j := 0; j < 9; j++ {
			if board[j][i] >= '1' && board[j][i] <= '9' {
				if m[board[j][i]-'1'] == 0 {
					m[board[j][i]-'1']++
				} else {
					return false
				}
			}
		}
	}
	return true
}

func cell(board [][]byte) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			m := [9]int{}
			for a := 0; a < 3; a++ {
				for b := 0; b < 3; b++ {
					if board[i+a][j+b] >= '1' && board[i+a][j+b] <= '9' {
						if m[board[i+a][j+b]-'1'] == 0 {
							m[board[i+a][j+b]-'1']++
						} else {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
