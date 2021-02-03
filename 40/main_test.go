package combinationsum2

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	type testCase struct {
		candidates []int
		target     int
		want       [][]int
	}
	testGroup := map[string]testCase{
		"case1": {
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			},
		},
		"case2": {
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := CombinationSum2(tC.candidates, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
