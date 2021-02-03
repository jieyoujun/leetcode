package dividetwointegers

import (
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		dividend int
		divisor  int
		want     int
	}
	testGroup := map[string]testCase{
		"case1": testCase{10, 3, 3},
		"case2": testCase{7, -3, -2},
		"case3": testCase{1, 1, 1},
		"case4": testCase{-2147483648, -1, 2147483647},
		"case5": testCase{-2147483648, 1, -2147483648},
		"case6": testCase{2147483647, -1, -2147483647},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Divide(tC.dividend, tC.divisor)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
