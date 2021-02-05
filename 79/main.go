package wordsearch

func exist(board [][]byte, word string) bool {
	// m := map[string]bool // 不需要做map, 太大了

	m, n := len(board), len(board[0])
	// 先定位一个字符
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				// 然后上下左右找后续字符
				board[i][j] = 0
				if search(board, word[1:], i, j) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func search(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	// 上下左右
	if i > 0 && board[i-1][j] == word[0] {
		board[i-1][j] = 0
		if search(board, word[1:], i-1, j) {
			return true
		}
		board[i-1][j] = word[0]
	}
	if i < len(board)-1 && board[i+1][j] == word[0] {
		board[i+1][j] = 0
		if search(board, word[1:], i+1, j) {
			return true
		}
		board[i+1][j] = word[0]
	}
	if j > 0 && board[i][j-1] == word[0] {
		board[i][j-1] = 0
		if search(board, word[1:], i, j-1) {
			return true
		}
		board[i][j-1] = word[0]
	}
	if j < len(board[0])-1 && board[i][j+1] == word[0] {
		board[i][j+1] = 0
		if search(board, word[1:], i, j+1) {
			return true
		}
		board[i][j+1] = word[0]
	}
	return false
}
