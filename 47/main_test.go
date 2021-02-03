package permutationsii

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	type testCase struct {
		nums []int
		want [][]int
	}
	testGroup := map[string]testCase{
		"case1": {
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		"case2": {
			[]int{1, 1, 2},
			[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := PermuteUnique(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
