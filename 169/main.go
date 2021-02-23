package majorityelement

func majorityElement(nums []int) int {
	// var res []int
	// for i := 0; i < len(nums); i++ {
	// 	if len(res) == 0 || nums[i] == res[0] {
	// 		res = append(res, nums[i])
	// 	} else {
	// 		res = res[:len(res)-1]
	// 	}
	// 	if len(res) > len(nums)/2 {
	// 		break
	// 	}
	// }
	// return res[0]

	// 优化空间
	res, count, n := nums[0], 1, len(nums)
	for i := 1; i < n; i++ {
		if count == 0 || nums[i] == res {
			res = nums[i]
			count++
		} else {
			count--
		}
		if count > n/2 {
			break
		}
	}
	return res
}
