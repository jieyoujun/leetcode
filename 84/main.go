package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	indexOfMinH := 0
	for i := 1; i < len(heights); i++ {
		if heights[indexOfMinH] > heights[i] {
			indexOfMinH = i
		}
	}
	return max(heights[indexOfMinH]*len(heights), largestRectangleArea(heights[:indexOfMinH]), largestRectangleArea(heights[indexOfMinH+1:]))
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
