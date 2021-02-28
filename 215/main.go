package kthlargestelementinanarray

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]

	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	// return nums[len(nums)-k]
}
