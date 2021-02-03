package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 1, 2}, 2},
		"case2": testCase{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		"case3": testCase{[]int{0, 0, 0, 0, 0, 0, 0, 0}, 1},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := RemoveDuplicates(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
