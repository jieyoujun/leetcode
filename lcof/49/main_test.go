package choushulcof

import (
	"reflect"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	type testCase struct {
		n    int
		want int
	}
	testGroup := map[string]testCase{
		"case1": {1, 1},
		"case2": {10, 12},
		"case3": {17, 27},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := NthUglyNumber(tC.n)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
