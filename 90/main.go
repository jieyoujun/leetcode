package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	dfs(nums, 0, nil, &results)
	return results
}

func dfs(nums []int, left int, result []int, results *[][]int) {
	*results = append(*results, append([]int{}, result...))
	for i := left; i < len(nums); i++ {
		if i > left && nums[i] == nums[i-1] {
			continue
		}
		result = append(result, nums[i])
		dfs(nums, i+1, result, results)
		result = result[:len(result)-1]
	}
}
