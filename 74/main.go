package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 二分查找定位行
	up, down := 0, m-1
	for up <= down {
		middleRow := (up + down) / 2
		if target > matrix[middleRow][n-1] {
			up = middleRow + 1
		} else if target < matrix[middleRow][0] {
			down = middleRow - 1
		} else {
			// 二分查找定位列
			left, right := 0, n-1
			for left <= right {
				middleCol := (left + right) / 2
				if target > matrix[middleRow][middleCol] {
					left = middleCol + 1
				} else if target < matrix[middleRow][middleCol] {
					right = middleCol - 1
				} else {
					return true
				}
			}
			break
		}
	}
	return false
}
