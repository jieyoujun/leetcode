package erchasousuoshudehouxubianlixulielcof

import (
	"reflect"
	"testing"
)

func TestVerifyPostorder(t *testing.T) {
	type testCase struct {
		postorder []int
		want      bool
	}
	testGroup := map[string]testCase{
		"case1": {
			[]int{1, 6, 3, 2, 5},
			false,
		},
		"case2": {
			[]int{5, 7, 6, 9, 11, 10, 8},
			true,
		},
		"case3": {
			[]int{7, 4, 6, 5},
			false,
		},
		"case4": {
			[]int{1, 3, 2, 6, 5},
			true,
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := VerifyPostorder(tC.postorder)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
