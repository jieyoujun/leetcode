package shunshizhendayinjuzhenlcof

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	type testCase struct {
		matrix [][]int
		want   []int
	}
	testGroup := map[string]testCase{
		"case1": {
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		"case2": {
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := SpiralOrder(tC.matrix)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
