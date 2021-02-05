package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) (int, []int) {
	num, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			if count == 2 {
				// 不要这个nums[i]
				nums = append(nums[:i], nums[i+1:]...)
				i-- // 删了之后i当然得--
				continue
			}
			count++
		} else {
			num = nums[i]
			count = 1
		}
	}
	return len(nums), nums
}
