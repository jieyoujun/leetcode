package removeduplicatesfromsortedarray

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] != nums[i] {
			// 只记不一样的
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
