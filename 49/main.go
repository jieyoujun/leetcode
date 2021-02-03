package groupanagrams

func GroupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	// 排序
	// m := make(map[string][]string)
	// for _, str := range strs {
	// 	tmp := []rune(str)
	// 	sort.Slice(tmp, func(i, j int) bool {
	// 		return tmp[i] < tmp[j]
	// 	})
	// 	m[string(tmp)] = append(m[string(tmp)], str)
	// }

	// 计数
	m := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := [26]int{}
		for _, r := range str {
			tmp[r-'a']++ // 所有输入均为小写字母。
		}
		m[tmp] = append(m[tmp], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
