package plusone

func plusOne(digits []int) []int {
	// 最后一个数加1
	// 从后往前遍历digits
	// digit += carried
	// carried = digit / 10
	// digit %= 10
	digits[len(digits)-1]++
	var carried int
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carried
		carried = digits[i] / 10
		if carried == 0 {
			break
		}
		digits[i] %= 10
	}
	if carried == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
