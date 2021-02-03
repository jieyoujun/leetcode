package countandsay

import (
	"strconv"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	say := CountAndSay(n - 1)
	res := ""
	for i := 0; i < len(say); i++ {
		count := 1
		for i < len(say)-1 && say[i] == say[i+1] {
			count++
			i++
		}
		res += strconv.Itoa(count)
		res += string(say[i])
	}
	return res
}
