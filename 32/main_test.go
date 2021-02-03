package longestvalidparentheses

import (
	"reflect"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	type testCase struct {
		s    string
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{"(()", 2},
		"case2": testCase{")()())", 4},
		"case3": testCase{"(", 0},
		"case4": testCase{"()(())", 6},
		"case5": testCase{"()((()))", 8},
		"case6": testCase{"((()))())", 8},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := LongestValidParentheses(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
