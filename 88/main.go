package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for n > 0 {
		for m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1], nums1[m-1] = nums1[m-1], nums1[m+n-1]
			m--
		}
		nums1[m+n-1], nums2[n-1] = nums2[n-1], nums1[m+n-1]
		n--
	}
	return nums1
}
