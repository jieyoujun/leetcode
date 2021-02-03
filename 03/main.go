package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	left, max := 0, 0
	m := make(map[byte]int)
	for right := range s {
		if _, ok := m[s[right]]; ok {
			if left < m[s[right]] {
				left = m[s[right]]
			}
		}
		if max < right-left+1 {
			max = right - left + 1
		}
		m[s[right]] = right + 1
	}
	return max
}
