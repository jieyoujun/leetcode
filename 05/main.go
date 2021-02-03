package longestpalindrome

import (
	"strings"
)

// 错误思路: 没有考虑到偶数分布abba
func LongestPalindrome(s string) string {
	var ans string

	// 把偶数转换成奇数情况
	sTmp := "#"
	for _, r := range s {
		sTmp += string(r)
		sTmp += "#"
	}
	s = sTmp

	maxStep, step := 0, 0
	for i := 0; i < len(s); i++ {
		step = 0
		for {
			if i-step <= 0 || i+step >= len(s)-1 {
				break
			}
			step++
			if s[i-step] != s[i+step] {
				step--
				break
			}
		}
		if maxStep <= step {
			ans = s[i-step : i+step+1]
			maxStep = step
		}
	}
	return strings.ReplaceAll(ans, "#", "")
}
