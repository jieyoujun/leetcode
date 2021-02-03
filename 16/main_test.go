package threesumclosest

import (
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{-1, 2, 1, -4}, 1, 2},
		"case2": testCase{[]int{1, 1, 1, 1}, 0, 3},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := ThreeSumClosest(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
