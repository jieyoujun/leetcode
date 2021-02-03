package maximumsubarray

func MaxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	max, tmp := nums[0], 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if max < tmp {
			max = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return max
}
