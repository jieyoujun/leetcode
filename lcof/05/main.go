package tihuankonggelcof

func ReplaceSpace(s string) string {
	// return strings.ReplaceAll(s, " ", "%20")
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res += "%20"
		} else {
			res += string(s[i])
		}
	}
	return res
}
