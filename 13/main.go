package romantointeger

func RomanToInt(s string) int {
	var res int
	if len(s) < 1 {
		return res
	}
	r2I := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s)-1; i++ {
		if r2I[s[i]] >= r2I[s[i+1]] {
			res += r2I[s[i]]
		} else {
			res -= r2I[s[i]]
		}
	}
	return res + r2I[s[len(s)-1]]
}
