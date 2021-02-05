package combinations

func combine(n int, k int) [][]int {
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var results [][]int
	if len(nums) > 0 {
		dfs(nums, k, nil, &results)
	}
	return results
}

func dfs(nums []int, k int, result []int, results *[][]int) {
	if len(result) == k {
		*results = append(*results, append([]int{}, result...))
		return
	}
	for idx, num := range nums {
		dfs(nums[idx+1:], k, append(result, num), results)
	}
}
