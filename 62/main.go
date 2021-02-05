package uniquepaths

func uniquePaths(m int, n int) int {
	// m = 3, n = 4
	// 1 1 1 1 为什么这一行都是1呢, 因为只能向右向下走
	// 1 2 3 4
	// 1 3 6 10
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + // 向下走过来
				dp[i][j-1] // 向右走过来
		}
	}
	return dp[m-1][n-1]

	// res := 0
	// dfs(m, n, 1, 1, &res)
	// return res
}

// func dfs(m, n, i, j int, res *int) {
// 	if i+j == m+n {
// 		if i == m && j == n {
// 			*res++
// 		}
// 		return
// 	}
// 	if i < m {
// 		dfs(m, n, i+1, j, res)
// 	}
// 	if j < n {
// 		dfs(m, n, i, j+1, res)
// 	}
// }
