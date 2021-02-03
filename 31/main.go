package nextpermutation

func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	reverse(nums, i, len(nums)-1)
	if i == 0 {
		return
	}
	j := i - 1
	for i < len(nums) && nums[j] >= nums[i] {
		i++
	}
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func reverse(nums []int, start, end int) {
	var tmp int
	for start < end {
		tmp = nums[end]
		nums[end] = nums[start]
		nums[start] = tmp
		start++
		end--
	}
}
