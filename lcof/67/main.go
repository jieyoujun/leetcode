package bazifuchuanzhuanhuanchengzhengshulcof

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	res, sign := 0, 1

	for idx, num := range str {
		if num >= '0' && num <= '9' {
			res = res*10 + int(num-'0')
		} else if idx == 0 && num == '-' {
			sign = -1
		} else if idx == 0 && num == '+' {
			sign = 1
		} else {
			break
		}

		if res > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return res * sign
}
