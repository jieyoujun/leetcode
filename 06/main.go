package zigzagconversion

import "strings"

// https://www.codenong.com/cs106905963/
func Convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	indices := []int{0}
	for i := 1; i < numRows; i++ {
		indices = append(indices, i)
	}
	for i := numRows - 2; i > 0; i-- {
		indices = append(indices, i)
	}
	ans := make([]string, numRows)
	for idx, r := range s {
		ans[indices[idx%len(indices)]] += string(r)
	}
	return strings.Join(ans, "")
}

// func Convert(s string, numRows int) string {
// 	if len(s) <= numRows || numRows == 1 {
// 		return s
// 	}
// 	var ans string
// 	for i := 0; i < numRows; i++ {
// 		for idx, r := range s {
// 			if idx%(2*numRows-2) == i || idx%(2*numRows-2) == (2*numRows-2-i) {
// 				ans += string(r)
// 			}
// 		}
// 	}
// 	return ans
// }

// i%(2n-2) == 0 -> row_0
// i%(2n-2) == 1 || i%(2n-2) == (2n-2-1) -> row_1
// ...

// for idx, r := range s {
// 	if idx%(2*numRows-2) == 0 {
// 		ans += string(r)
// 	}
// }
// for idx, r := range s {
// 	if idx%(2*numRows-2) == 1 || idx%(2*numRows-2) == (2*numRows-2-1) {
// 		ans += string(r)
// 	}
// }
