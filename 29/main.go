package dividetwointegers

import (
	"math"
)

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := (dividend ^ divisor) < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := 0
	for i := 31; i >= 0; i-- {
		// dividend >> i 等价于 dividend / 2^n
		if (dividend >> i) >= divisor {
			res += 1 << i
			dividend -= divisor << i
		}
	}
	if neg {
		res = -res
	}
	return res
}
