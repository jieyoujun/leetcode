package numberofislands

func numIslands(grid [][]byte) int {
	nIsland := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				nIsland++
				dfs(grid, i, j) // 把周围1都置0
			}
		}
	}
	return nIsland
}

func dfs(grid [][]byte, i, j int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
}
