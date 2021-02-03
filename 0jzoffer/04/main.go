package erweishuzuzhongdechazhaolcof

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	res := false
	if matrix != nil && len(matrix) > 0 && len(matrix[0]) > 0 {
		row, col := 0, len(matrix[0])-1
		for row <= len(matrix)-1 && col >= 0 {
			if matrix[row][col] == target {
				res = true
				break
			} else if matrix[row][col] > target {
				col--
			} else {
				row++
			}
		}
	}
	return res
}
