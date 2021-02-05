package decodeways

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	ns := len(s)
	dp := make([]int, ns+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= ns; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[ns]
}
