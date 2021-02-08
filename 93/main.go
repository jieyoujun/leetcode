package restoreipaddresses

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var results []string
	dfs(s, nil, &results)
	return results
}

func dfs(s string, result []string, results *[]string) {
	if len(s)+len(result)*3 > 12 {
		return
	}
	if len(s) > 0 && len(result)*3 == 9 {
		if (len(s) > 1 && s[0] == '0') || (len(s) == 3 && strings.Compare("255", s) < 0) {
			return
		}
		result = append(result, s)
		*results = append(*results, strings.Join(result, "."))
		return
	}
	for i := 1; i <= 3; i++ {
		if len(s) >= i {
			if i == 3 && strings.Compare("255", s[:i]) < 0 {
				break
			}
			result = append(result, s[:i])
			dfs(s[i:], result, results)
			result = result[:len(result)-1]
		}
		if len(s) > 0 && s[0] == '0' {
			break
		}
	}
}
