package longestsubstringwithoutrepeatingcharacters

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type testCase struct {
		s    string
		want int
	}
	testGroup := map[string]testCase{
		"case1": {"abcabcbb", 3},
		"case2": {"bbbbb", 1},
		"case3": {"pwwkew", 3},
		"case4": {" ", 1},
		"case5": {"au", 2},
		"case6": {"arabcacfr", 4},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := LengthOfLongestSubstring(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
