package twosum

func TwoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	m := make(map[int][]int, len(nums))
	for idx, v := range nums {
		m[v] = append(m[v], idx)
	}
	for idx, v := range nums {
		if _, ok := m[target-v]; ok {
			if len(m[target-v]) == 2 {
				return m[target-v]
			}
			if m[target-v][0] != idx {
				return []int{idx, m[target-v][0]}
			}
		}
	}
	return []int{}
}
