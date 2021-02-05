package searchinrotatedsortedarrayii

import "sort"

func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if sort.IntsAreSorted(nums[mid:]) {
			// 右半段有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 左半段有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
