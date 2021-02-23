package shuzuzhongdenixuduilcof

func ReversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return merge(nums, make([]int, len(nums)), 0, len(nums)-1)
}

func merge(nums, copy []int, left, right int) int {
	if left == right {
		return 0
	}
	mid := (left + right) / 2

	leftCount := merge(nums, copy, left, mid)
	rightCount := merge(nums, copy, mid+1, right)
	mergeCount := 0

	index := left
	i, j := left, mid+1
	for i <= mid && j <= right {
		// 反着比? 负负得正?
		if nums[i] > nums[j] {
			mergeCount += mid - i + 1
			copy[index] = nums[j] // 入辅助数组
			j++
		} else {
			copy[index] = nums[i]
			i++
		}
		index++
	}
	for i <= mid {
		copy[index] = nums[i]
		i++
		index++
	}
	for j <= right {
		copy[index] = nums[j]
		j++
		index++
	}
	for i := left; i <= right; i++ {
		nums[i] = copy[i]
	}
	return leftCount + rightCount + mergeCount
}
