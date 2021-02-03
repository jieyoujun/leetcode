package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, curInterval := range intervals {
		if curInterval[0] > right {
			// curInterval on right
			if !merged {
				res = append(res, []int{left, right})
				merged = true
			}
			res = append(res, curInterval)
		} else if curInterval[1] < left {
			// curInterval on left
			res = append(res, curInterval)
		} else {
			if left > curInterval[0] {
				left = curInterval[0]
			}
			if right < curInterval[1] {
				right = curInterval[1]
			}
		}
	}
	if !merged {
		res = append(res, []int{left, right})
	}
	return res
}
