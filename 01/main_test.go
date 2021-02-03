package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"case2": testCase{[]int{3, 2, 4}, 6, []int{1, 2}},
		"case3": testCase{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
		"case4": testCase{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := TwoSum(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
