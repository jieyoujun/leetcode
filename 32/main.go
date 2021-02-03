package longestvalidparentheses

func LongestValidParentheses(s string) int {
	res, dp := 0, make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			j := i - dp[i-1] - 1
			if j >= 0 && s[j] == '(' {
				if (j - 1) >= 0 {
					dp[i] = i - j + 1 + dp[j-1]
				} else {
					dp[i] = i - j + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
