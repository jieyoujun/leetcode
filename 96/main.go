package uniquebinarysearchtrees

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 递归
// func numTrees(n int) int {
// 	if n < 1 {
// 		return 1
// 	}
// 	return trees(1, n)
// }

// func trees(left, right int) int {
// 	t := 0
// 	if left > right {
// 		t++
// 		return t
// 	}
// 	for i := left; i <= right; i++ {
// 		t += trees(left, i-1) * trees(i+1, right)
// 	}
// 	return t
// }
