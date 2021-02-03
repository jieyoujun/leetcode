package stringtointegeratoi

import (
	"fmt"
	"math"
	"strings"
)

func MyAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) < 1 {
		return 0
	}
	str0 := str[0]
	neg := false
	if str0 == '-' || str0 == '+' {
		if str0 == '-' {
			neg = true
		}
		str = str[1:]
	} else if str0 -= '0'; str0 < 0 || str0 > 9 {
		return 0
	}
	n, sLen := 0, -1
	for _, r := range str {
		r -= '0'
		if r != 0 && sLen < 0 {
			sLen = 0
		}
		if r > 9 || r < 0 {
			break
		}
		n = n*10 + int(r)
		if sLen >= 0 {
			sLen++
		}
	}
	fmt.Println(str, neg, sLen, n)
	if sLen <= 10 && n <= math.MaxInt32 {
		if neg {
			return -n
		}
		return n
	}
	if neg {
		return math.MinInt32
	}
	return math.MaxInt32
}
