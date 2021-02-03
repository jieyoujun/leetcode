package searchinsertposition

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   int
	}
	testGroup := map[string]testCase{
		"case1": {[]int{1, 3, 5, 6}, 5, 2},
		"case2": {[]int{1, 3, 5, 6}, 2, 1},
		"case3": {[]int{1, 3, 5, 6}, 7, 4},
		"case4": {[]int{1, 3, 5, 6}, 0, 0},
		"case5": {[]int{}, 1, 0},
		"case6": {[]int{1}, 0, 0},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := SearchInsert(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
