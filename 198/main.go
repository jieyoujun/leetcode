package houserobber

func rob(nums []int) int {
	// 动态规划
	// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) >= 2 {
		nums[1] = max(nums[0], nums[1])
	}
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
	// dfs
	// var max int
	// for i := 0; i < len(nums); i++ {
	// 	dfs(nums, i+2, []int{nums[i]}, &max)
	// }
	// return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func dfs(nums []int, left int, money []int, max *int) {
// 	n := len(nums)
// 	if left+1 >= n {
// 		sum := 0
// 		for _, m := range money {
// 			sum += m
// 		}
// 		if *max < sum {
// 			*max = sum
// 		}
// 	}
// 	for i := left; i < n; i++ {
// 		dfs(nums, i+2, append(money, nums[i]), max)
// 	}
// }
