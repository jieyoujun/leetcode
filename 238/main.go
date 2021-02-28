package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefixProduct, suffixProduct := make([]int, n), make([]int, n)
	prefixProduct[0] = 1
	suffixProduct[n-1] = 1
	for i := 1; i < n; i++ {
		prefixProduct[i] = nums[i-1] * prefixProduct[i-1]
		suffixProduct[n-1-i] = nums[n-1-i+1] * suffixProduct[n-1-i+1]
	}
	for i := 0; i < n; i++ {
		prefixProduct[i] *= suffixProduct[i]
	}
	return prefixProduct
	// O(n*n)超时啦
	// result, total := []int{}, 0
	// for i := 0; i < len(nums); i++ {
	// 	result = append(result, mulitply(nums, i))
	// }
	// return result
}

func mulitply(nums []int, exclude int) int {
	result := 1
	for idx, num := range nums {
		if idx == exclude {
			continue
		}
		if num == 0 {
			return 0
		}
		result *= num
	}
	return result
}
