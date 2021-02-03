package combinationsum

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	var result [][]int
	sort.Ints(candidates)
	dfs(candidates, nil, target, 0, &result)
	return result
}

func dfs(candidates []int, nums []int, target, left int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}

	for i := left; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, result)
	}
}
