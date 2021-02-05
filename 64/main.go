package minimumpathsum

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := range grid {
		if i > 0 {
			grid[i][0] += grid[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j-1] < grid[i-1][j] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[m-1][n-1]
}
