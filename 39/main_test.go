package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type testCase struct {
		candidates []int
		target     int
		want       [][]int
	}
	testGroup := map[string]testCase{
		"case1": {
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{7},
				{2, 2, 3},
			},
		},
		"case2": {
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := CombinationSum(tC.candidates, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
