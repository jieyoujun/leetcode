package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}
	testGroup := map[string]testCase{
		"case1": {[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		"case2": {[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		"case3": {[]int{1}, 1, []int{0, 0}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := SearchRange(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
