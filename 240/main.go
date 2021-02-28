package searcha2dmatrixii

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	// 从右上角折线
	i, j := 0, n-1
	for i < m && j > 0 {
		if matrix[i][j] < target { // 往下
			i++
		} else if matrix[i][j] > target { // 往左
			j--
		} else {
			return true
		}
	}
	return false

	// 按行二分
	// for i := 0; i < m; i++ {
	// 	left, right := 0, n-1
	// 	for left <= right {
	// 		middle := (left + right) / 2
	// 		if target < matrix[i][middle] {
	// 			right = middle - 1
	// 		} else if target > matrix[i][middle] {
	// 			left = middle + 1
	// 		} else {
	// 			return true
	// 		}
	// 	}
	// }
	// return false
}
