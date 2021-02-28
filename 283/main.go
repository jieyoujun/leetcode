package movezeroes

func moveZeroes(nums []int) []int {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
	return nums

	// 空间换时间
	// nzeros := []int{}
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != 0 {
	// 		nzeros = append(nzeros, nums[i])
	// 	}
	// }
	// for i := 0; i < len(nums); i++ {
	// 	if i < len(nzeros) {
	// 		nums[i] = nzeros[i]
	// 	} else {
	// 		nums[i] = 0
	// 	}
	// }
	// return nums
}
