package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 位图
	// var res [][]int
	// max := 0
	// bitMap := make([]bool, 20000+1)
	// for _, interval := range intervals {
	// 	if max < interval[1]*2+1 {
	// 		max = interval[1]*2 + 1
	// 	}
	// 	for i := interval[0] * 2; i < interval[1]*2+1; i++ {
	// 		bitMap[i] = true
	// 	}
	// }

	// for index := 0; index <= max; {
	// 	start := nextTrue(bitMap, index)
	// 	if start < 0 {
	// 		break
	// 	}
	// 	end := nextFalse(bitMap, start)
	// 	res = append(res, []int{start / 2, (end - 1) / 2})
	// 	index = end
	// }
	// return res

	// 排序+合并
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	var (
		lastInterval = intervals[0]
		res          = [][]int{lastInterval}
	)
	for _, curInterval := range intervals[1:] {
		// end1 < start2
		if len(res) == 0 || lastInterval[1] < curInterval[0] {
			res = append(res, curInterval)
			lastInterval = curInterval
		}
		// end1 >=start2 && end1 < end2
		if lastInterval[1] < curInterval[1] {
			lastInterval[1] = curInterval[1]
		}
	}
	return res
}

func nextTrue(bitMap []bool, fromIndex int) int {
	for {
		if fromIndex >= len(bitMap)-1 {
			return -1
		}
		if bitMap[fromIndex] {
			return fromIndex
		}
		fromIndex++
	}
}

func nextFalse(bitMap []bool, fromIndex int) int {
	for {
		if fromIndex >= len(bitMap)-1 {
			return -1
		}
		if !bitMap[fromIndex] {
			return fromIndex
		}
		fromIndex++
	}
}
