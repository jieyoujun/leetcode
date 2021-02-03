package threesum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	left := 0
	for left < len(nums)-2 {
		if left != 0 && nums[left] == nums[left-1] {
			left++
			continue
		}
		middle := left + 1
		right := len(nums) - 1
		for middle < right {
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if middle > left+1 && nums[middle] == nums[middle-1] {
				middle++
				continue
			}
			if nums[left]+nums[middle]+nums[right] > 0 {
				right--
			} else if nums[left]+nums[middle]+nums[right] < 0 {
				middle++
			} else {
				res = append(res, []int{nums[left], nums[middle], nums[right]})
				right--
				middle++
			}
		}
		left++
	}
	return res
}
