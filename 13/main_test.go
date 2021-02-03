package romantointeger

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type testCase struct {
		s    string
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{"III", 3},
		"case2": testCase{"IV", 4},
		"case3": testCase{"IX", 9},
		"case4": testCase{"LVIII", 58},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := RomanToInt(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
