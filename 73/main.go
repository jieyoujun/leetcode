package setmatrixzeroes

func setZeroes(matrix [][]int) [][]int {
	if matrix != nil {
		m, n := len(matrix), len(matrix[0])
		r0, c0 := false, false
		for j := 0; j < n; j++ {
			if matrix[0][j] == 0 { // 第一行有0
				r0 = true
				break
			}
		}
		for i := 0; i < m; i++ {
			if matrix[i][0] == 0 { // 第一列有0
				c0 = true
				break
			}
		}

		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				if matrix[i][j] == 0 {
					// 把第一行和第一列作为标志位
					matrix[i][0] = 0
					matrix[0][j] = 0
				}
			}
		}

		// 置零
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				if matrix[i][0] == 0 || matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}
		}

		if r0 {
			for j := 0; j < n; j++ {
				matrix[0][j] = 0
			}
		}
		if c0 {
			for i := 0; i < m; i++ {
				matrix[i][0] = 0
			}
		}
	}
	return matrix

	// if matrix != nil {
	// 	m, n := len(matrix), len(matrix[0])
	// 	var r0, c0 []int
	// 	for i := 0; i < m; i++ {
	// 		for j := 0; j < n; j++ {
	// 			if matrix[i][j] == 0 {
	// 				r0 = append(r0, i)
	// 				c0 = append(c0, j)
	// 			}
	// 		}
	// 	}
	// 	for _, r := range r0 {
	// 		matrix[r] = make([]int, n)
	// 	}
	// 	for i := 0; i < m; i++ {
	// 		for _, c := range c0 {
	// 			matrix[i][c] = 0
	// 		}
	// 	}
	// }
	// return matrix
}
