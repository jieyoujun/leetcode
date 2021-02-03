package validparentheses

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	type testCase struct {
		s    string
		want bool
	}
	testGroup := map[string]testCase{
		"case1": testCase{"()", true},
		"case2": testCase{"([]{})", true},
		"case3": testCase{"([{}])", true},
		"case4": testCase{"({)}", false},
		"case5": testCase{"({})()", true},
		"case6": testCase{"{}{}}", false},
		"case7": testCase{"}(", false},
		"case8": testCase{"({[)", false},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := IsValid(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
