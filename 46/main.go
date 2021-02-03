package permutations

func Permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	var result [][]int

	// 递归
	// for i, num := range nums {
	// 	// 把num从nums去除得到tmp
	// 	tmp := make([]int, len(nums)-1)
	// 	copy(tmp[0:], nums[:i])
	// 	copy(tmp[i:], nums[i+1:])

	// 	// 对tmp全排列
	// 	tmpPermuted := Permute(tmp)
	// 	for _, tp := range tmpPermuted {
	// 		result = append(result, append(tp, num))
	// 	}
	// }

	// dfs
	dfs(nums, 0, &result)
	return result
}

func dfs(nums []int, left int, result *[][]int) {
	if left == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
	}

	for i := left; i < len(nums); i++ {
		// nums[:left]已定, nums[left:]都是nums[left]的待选
		nums[i], nums[left] = nums[left], nums[i]
		dfs(nums, left+1, result)
	}
	// 还原
	for i := len(nums) - 1; i > left; i-- {
		nums[i], nums[left] = nums[left], nums[i]
	}
}
