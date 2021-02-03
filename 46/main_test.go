package permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
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
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Permute(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
