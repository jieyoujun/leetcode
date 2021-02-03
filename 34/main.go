package findfirstandlastpositionofelementinsortedarray

func SearchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			midL, midR := mid, mid
			for midL-1 >= left && nums[midL-1] == target {
				midL--
			}
			for midR+1 <= right && nums[midR+1] == target {
				midR++
			}
			return []int{midL, midR}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
