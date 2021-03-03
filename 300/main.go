package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([]int, n)
	dp[0] = 1
	max := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max

	// dfs超时
	// max := 1
	// for i := 0; i < n; i++ {
	// 	dfs(nums, i+1, []int{nums[i]}, &max)
	// }
	// return max
}

func dfs(nums []int, left int, seq []int, max *int) {
	nNums, nSeq := len(nums), len(seq)
	if *max < nSeq {
		*max = nSeq
	}
	if left == nNums {
		return
	}
	for i := left; i < nNums; i++ {
		if nums[i] > seq[nSeq-1] {
			dfs(nums, i+1, append(seq, nums[i]), max)
		}
	}
}
