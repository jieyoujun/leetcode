package lengthoflastword

func lengthOfLastWord(s string) int {
	for {
		if len(s) == 0 {
			return 0
		}
		if s[len(s)-1] != ' ' {
			break
		}
		s = s[:len(s)-1]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s) - 1 - i
		}
		if i == 0 {
			return len(s)
		}
	}
	return 0
}
