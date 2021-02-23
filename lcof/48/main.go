package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	if len(s) > 0 {
		curLen, maxLen, preIndices := 0, 0, [256]int{}
		// ascii有256个字符
		for i := 0; i < 256; i++ {
			preIndices[i] = -1
		}
		for i := 0; i < len(s); i++ {
			if preIndices[s[i]] < 0 || i-preIndices[s[i]] > curLen {
				// 没出现过 或者 d > f(i-1) => f(i) = f(i-1)+1
				curLen++
			} else {
				if maxLen < curLen {
					maxLen = curLen
				}
				// f(i) = d
				curLen = i - preIndices[s[i]]
			}
			preIndices[s[i]] = i
		}
		if maxLen < curLen {
			maxLen = curLen
		}
		return maxLen
	}
	return 0
}
