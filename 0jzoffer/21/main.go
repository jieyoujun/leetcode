package diaozhengshuzushunxushiqishuweiyuoushuqianmianlcof

func Exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)-1
	var tmp int
	for left < right {
		for left < right && nums[left]&1 != 0 {
			left++
		}
		for left < right && nums[right]&1 == 0 {
			right--
		}
		if left < right {
			tmp = nums[left]
			nums[left] = nums[right]
			nums[right] = tmp
		}
	}
	return nums
}
