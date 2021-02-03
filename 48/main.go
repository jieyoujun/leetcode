package rotateimage

func Rotate(matrix [][]int) [][]int {
	// // 从左到右横着读
	// // 从下往上竖着读
	// if matrix == nil || (len(matrix) != len(matrix[0])) {
	// 	return nil
	// }
	// n := len(matrix)
	// var tmp [][]int
	// for i := 0; i < n; i++ {
	// 	tmp = append(tmp, make([]int, n))
	// }
	// for col := 0; col < n; col++ {
	// 	for row := n - 1; row >= 0; row-- {
	// 		tmp[col][n-row-1] = matrix[row][col]
	// 		// [row][col]->[n-row-1][col]->[col][n-row-1]
	// 		fmt.Print(matrix[row][col], " ")
	// 	}
	// 	fmt.Println()
	// }
	// copy(matrix, tmp)
	// return matrix

	n := len(matrix)
	for row := 0; row < n/2; row++ {
		matrix[row], matrix[n-row-1] = matrix[n-row-1], matrix[row]
	}
	for row := 0; row < n; row++ {
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
	return matrix
}
