package shuzuzhongchuxiancishuchaoguoyibandeshuzilcof

func majorityElement(nums []int) int {
	n, count := nums[0], 1
	for _, num := range nums[1:] {
		if n == num {
			count++
		} else {
			count--
		}
		if count == 0 {
			n = num
			count = 1
		}
	}
	return n
}
