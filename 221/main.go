package maximalsquare

// matrix[i][j]=='1' => dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n, max := len(matrix), len(matrix[0]), 0

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func min(a, b, c int) int {
	if b < a {
		if b < c {
			return b
		}
		return c
	}
	if a < c {
		return a
	}
	return c
}
