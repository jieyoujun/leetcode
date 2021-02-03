package generateparentheses

func GenerateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	return backTrack(n, n, "", []string{})
}

func backTrack(left, right int, track string, res []string) []string {
	if left == 0 && right == 0 {
		return append(res, track)
	}
	if left > 0 {
		res = backTrack(left-1, right, track+"(", res)
	}
	if right > 0 && left < right {
		// left < right => 先有前括号才有后括号
		res = backTrack(left, right-1, track+")", res)
	}
	return res
}
