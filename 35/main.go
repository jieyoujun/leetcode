package searchinsertposition

func SearchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	if target < nums[left] {
		return left
	}
	if target > nums[right] {
		return right + 1
	}
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
