package maximumsubarray

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		"case2": testCase{[]int{1, -2, 3, 10, -4, 7, 2, -5}, 18},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MaxSubArray(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
