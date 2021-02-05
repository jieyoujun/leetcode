package sortcolors

func sortColors(nums []int) []int {
	// 0往前, 2往后
	i, left, right := 0, 0, len(nums)-1
	for i <= right {
		if nums[i] == 0 && i >= left {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 && i <= right {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
	return nums

	// 不要依赖代码库啦
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] <= nums[j]
	// })
	// return nums
}
