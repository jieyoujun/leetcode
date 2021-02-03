package nextpermutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type testCase struct {
		nums []int
		want []int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 2, 3}, []int{1, 3, 2}},
		"case2": testCase{[]int{3, 2, 1}, []int{1, 2, 3}},
		"case3": testCase{[]int{1, 1, 5}, []int{1, 5, 1}},
		"case4": testCase{[]int{1, 3, 2}, []int{2, 1, 3}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			NextPermutation(tC.nums)
			if !reflect.DeepEqual(tC.nums, tC.want) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, tC.nums)
			}
		})
	}
}
