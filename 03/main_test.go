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
		"case1": testCase{"abcabcbb", 3},
		"case2": testCase{"bbbbb", 1},
		"case3": testCase{"pwwkew", 3},
		"case4": testCase{" ", 1},
		"case5": testCase{"au", 2},
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
