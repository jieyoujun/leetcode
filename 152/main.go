package maximumproductsubarray

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	p1, p2, maxp := nums[0], nums[0], nums[0]
	for i := 0; i < n-1; i++ {
		if nums[i+1] > 0 {
			p1, p2 = max(p1*nums[i+1], nums[i+1]), min(p2*nums[i+1], nums[i+1])
		} else {
			p1, p2 = max(p2*nums[i+1], nums[i+1]), min(p1*nums[i+1], nums[i+1])
		}
		if maxp < p1 {
			maxp = p1
		}
	}
	return maxp
	// 暴力
	// maxp := nums[0]
	// for i := 0; i < n; i++ {
	// 	product := 1
	// 	for j := i; j < n; j++ {
	// 		product *= nums[j]
	// 		if maxp < product {
	// 			maxp = product
	// 		}
	// 	}
	// }
	// return maxp
}
