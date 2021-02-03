package containerwithmostwater

func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max, left, right := 0, 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if max < (right-left)*height[left] {
				max = (right - left) * height[left]
			}
			left++
		} else {
			if max < (right-left)*height[right] {
				max = (right - left) * height[right]
			}
			right--
		}
	}
	return max
}

// func MaxArea(height []int) int {
// 	hLen := len(height)
// 	if hLen < 2 {
// 		return 0
// 	}
// 	max := 0
// 	for i := 0; i < hLen-1; i++ {
// 		for j := hLen - 1; j > i; j-- {
// 			if height[i] < height[j] {
// 				if max < (j-i)*height[i] {
// 					max = (j - i) * height[i]
// 				}
// 			} else {
// 				if max < (j-i)*height[j] {
// 					max = (j - i) * height[j]
// 				}
// 			}
// 		}
// 	}
// 	return max
// }
