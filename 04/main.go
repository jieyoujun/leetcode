package medianoftwosortedarrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0, len(nums1)+len(nums2))
	// 先合并
	for {
		if len(nums1) == 0 {
			nums = append(nums, nums2...)
			break
		}
		if len(nums2) == 0 {
			nums = append(nums, nums1...)
			break
		}
		if nums1[0] < nums2[0] {
			nums = append(nums, nums1[0])
			nums1 = nums1[1:]
		} else {
			nums = append(nums, nums2[0])
			nums2 = nums2[1:]
		}
	}
	// 再找中位数
	sort.Ints(nums)
	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	}
	return float64((nums[len(nums)/2] + nums[len(nums)/2-1])) / 2
}
