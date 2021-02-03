package implementstrstr

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	type testCase struct {
		haystack string
		needle   string
		want     int
	}
	testGroup := map[string]testCase{
		"case1": testCase{"hello", "ll", 2},
		"case2": testCase{"aaaaa", "bba", -1},
		"case3": testCase{"aaa", "aaaa", -1},
		"case4": testCase{"mississippi", "issip", 4},
		"case5": testCase{"mississippi", "issipi", -1},
		"case6": testCase{"a", "a", 0},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := StrStr(tC.haystack, tC.needle)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
