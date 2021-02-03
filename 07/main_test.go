package reverseinteger

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type testCase struct {
		x    int
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{123, 321},
		"case2": testCase{-123, -321},
		"case3": testCase{120, 21},
		"case4": testCase{1534236469, 0},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Reverse(tC.x)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
