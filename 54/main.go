package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	// 从(0,0)开始向前(初始为右)走, 边界, 右转
	var visited [][]bool
	for row := 0; row < len(matrix); row++ {
		visited = append(visited, make([]bool, len(matrix[row])))
	}
	// fmt.Println(visited)
	m, n := len(matrix), len(matrix[0])
	i, j := 0, 0
	right, down, left, up := true, false, false, false
	var res []int
	for {
		if visited[i][j] == false {
			res = append(res, matrix[i][j])
			visited[i][j] = true
			if len(res) == m*n {
				break
			}
		}
		if right {
			if j == n-1 || visited[i][j+1] == true {
				right = false
				down = true
			} else {
				j++
			}
			continue
		}
		if down {
			if i == m-1 || visited[i+1][j] == true {
				down = false
				left = true
			} else {
				i++
			}
			continue
		}
		if left {
			if j == 0 || visited[i][j-1] == true {
				left = false
				up = true
			} else {
				j--
			}
			continue
		}
		if up {
			if i == 0 || visited[i-1][j] == true {
				up = false
				right = true
			} else {
				i--
			}
			continue
		}
	}
	return res
}
