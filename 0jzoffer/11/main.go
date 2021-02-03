package xuanzhuanshuzudezuixiaoshuzilcof

func MinArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if numbers[right] > numbers[mid] {
			// 右半有序, 最小在左边
			right = mid
		} else if numbers[right] < numbers[mid] {
			// 左半有序, 最小在右边
			left = mid + 1
		} else {
			// 去重
			right--
		}
	}
	return numbers[left]
}
