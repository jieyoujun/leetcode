package longestcommonprefix

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		strs []string
		want string
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]string{"flower", "flow", "flight"}, "fl"},
		"case2": testCase{[]string{"dog", "racecar", "car"}, ""},
		"case3": testCase{[]string{"aa", "a"}, "a"},
		"case4": testCase{[]string{"aa", ""}, ""},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := LongestCommonPrefix(tC.strs)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
