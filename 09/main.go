package palindromenumber

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev, x0 := 0, x
	for x != 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev == x0
}

// func IsPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	xStr := strconv.Itoa(x)
// 	xLen := len(xStr)
// 	for i := 0; i < xLen/2; i++ {
// 		if xStr[i] != xStr[xLen-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
