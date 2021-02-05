package addbinary

func addBinary(a string, b string) string {
	// 从后向前遍历a, b
	// a == 1 && b == 1 => sum = 0, carried = 1
	// a == 0 && b == 0 => sum = 0, carried = 1
	// else => sum = 1, carried = 0
	n := len(a)
	if n < len(b) {
		n = len(b)
	}
	res, carried := "", 0
	for i := n - 1; i >= 0; i-- {
		nexta, nextb := next(&a), next(&b)
		switch nexta + nextb + carried {
		case 3:
			res = "1" + res
			carried = 1
		case 2:
			res = "0" + res
			carried = 1
		case 1:
			res = "1" + res
			carried = 0
		case 0:
			res = "0" + res
			carried = 0
		}
	}
	if carried == 1 {
		res = "1" + res
	}
	return res
}

func next(s *string) (next int) {
	if len(*s) > 0 {
		if (*s)[len(*s)-1] == '1' {
			next = 1
		}
		*s = (*s)[:len(*s)-1]
	}
	return
}
