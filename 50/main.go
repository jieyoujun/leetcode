package powxn

func MyPow(x float64, n int) float64 {
	if x == float64(0.0) {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res float64 = 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}
