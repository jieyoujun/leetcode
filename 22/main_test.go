package generateparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type testCase struct {
		n    int
		want []string
	}
	testGroup := map[string]testCase{
		"case1": testCase{3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
		"case2": testCase{2, []string{
			"(())",
			"()()",
		}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := GenerateParenthesis(tC.n)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
