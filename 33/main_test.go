package searchinrotatedsortedarray

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		"case2": testCase{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Search(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
