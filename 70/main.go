package climbingstairs

func climbStairs(n int) int {
	return stj(n, 0, 1)
}

func stj(n, pre, cur int) int {
	if n == 0 {
		return cur
	}
	return stj(n-1, cur, pre+cur)
}
