package threesumclosest

import (
	"fmt"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 || len(nums) > 1000 {
		return 0
	}
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[len(nums)-1]
	left := 0
	for left < len(nums)-2 {
		if left > 0 && nums[left] == nums[left-1] {
			left++
			continue
		}
		middle := left + 1
		right := len(nums) - 1
		for middle < right {
			fmt.Println(left, middle, right, nums)
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if middle > left+1 && nums[middle] == nums[middle-1] {
				middle++
				continue
			}
			tmp := nums[left] + nums[middle] + nums[right]
			if (target-closest)*(target-closest) > (target-tmp)*(target-tmp) {
				closest = tmp
			}
			if tmp > target {
				right--
			} else if tmp > target {
				middle++
			} else {
				return target
			}
		}
		left++
	}
	return closest
}
