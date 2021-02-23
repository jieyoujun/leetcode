package shunshizhendayinjuzhenlcof

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) > 0 && len(matrix[0]) > 0 {
		res := []int{}
		for start := 0; len(matrix) > start*2 && len(matrix[0]) > start*2; start++ {
			printCircle(matrix, start, &res)
		}
		return res
	}
	return nil
}

func printCircle(matrix [][]int, start int, res *[]int) {
	endRow, endCol := len(matrix)-1-start, len(matrix[0])-1-start
	// 左上->右上
	for i := start; i <= endCol; i++ {
		*res = append(*res, matrix[start][i])
	}
	// 右上->右下
	if start < endRow {
		for i := start + 1; i <= endRow; i++ {
			*res = append(*res, matrix[i][endCol])
		}
	}
	// 右下->左下
	if start < endRow && start < endCol {
		for i := endCol - 1; i >= start; i-- {
			*res = append(*res, matrix[endRow][i])
		}
	}
	// 左下->左上
	if start < endRow-1 && start < endCol {
		for i := endRow - 1; i >= start+1; i-- {
			*res = append(*res, matrix[i][start])
		}
	}
}
