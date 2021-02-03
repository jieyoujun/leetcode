package reverseinteger

import (
	"math"
)

func Reverse(x int) int {
	var ans []int
	for x != 0 {
		ans = append(ans, x%10)
		x /= 10
	}
	var ret int
	for idx, v := range ans {
		ret += int(math.Pow10(len(ans)-1-idx)) * v
	}
	if ret <= math.MaxInt32 && ret >= math.MinInt32 {
		return ret
	}
	return 0
}
