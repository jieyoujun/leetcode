package foursum

import "sort"

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	for left := 0; left < len(nums)-3; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		if nums[left]+nums[left+1]+nums[left+2]+nums[left+3] > target {
			break
		}
		if nums[left]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1] < target {
			continue
		}
		for middle0 := left + 1; middle0 < len(nums)-2; middle0++ {
			if middle0 > left+1 && nums[middle0] == nums[middle0-1] {
				continue
			}
			if nums[left]+nums[middle0]+nums[middle0+1]+nums[middle0+2] > target {
				break
			}
			if nums[left]+nums[middle0]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				continue
			}
			middle1, right := middle0+1, len(nums)-1
			for middle1 < right {
				if middle1 > middle0+1 && nums[middle1] == nums[middle1-1] {
					middle1++
					continue
				}
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				if nums[left]+nums[middle0]+nums[middle1]+nums[right] == target {
					res = append(res, []int{nums[left], nums[middle0], nums[middle1], nums[right]})
					middle1++
					right--
				} else if nums[left]+nums[middle0]+nums[middle1]+nums[right] < target {
					middle1++
				} else {
					right--
				}
			}
		}
	}
	return res
}
