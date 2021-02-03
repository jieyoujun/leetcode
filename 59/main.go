package spiralmatrixii

func generateMatrix(n int) [][]int {

	var res [][]int
	// 初始化res
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	right, down, left, up := true, false, false, false
	i, j, num := 0, 0, 1
	for {
		if res[i][j] == 0 {
			res[i][j] = num
			if num == n*n {
				break
			}
			num++
		}
		if right {
			if j == n-1 || res[i][j+1] != 0 {
				right = false
				down = true
			} else {
				j++
			}
			continue
		}
		if down {
			if i == n-1 || res[i+1][j] != 0 {
				down = false
				left = true
			} else {
				i++
			}
			continue
		}
		if left {
			if j == 0 || res[i][j-1] != 0 {
				left = false
				up = true
			} else {
				j--
			}
			continue
		}
		if up {
			if i == 0 || res[i-1][j] != 0 {
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
