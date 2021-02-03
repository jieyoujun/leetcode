package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type testCase struct {
		nums []int
		want [][]int
	}
	testGroup := map[string]testCase{
		"case1": testCase{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
		"case2": testCase{
			[]int{0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
		"case3": testCase{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			[][]int{
				{0, 0, 0},
			},
		},
		"case4": testCase{
			[]int{0, -1, 1},
			[][]int{
				{-1, 0, 1},
			},
		},
		"case5": testCase{
			[]int{1, 1, -2},
			[][]int{
				{1, 1, -2},
			},
		},
		"case6": testCase{
			[]int{1, 1, 1},
			[][]int{},
		},
		"case7": testCase{
			[]int{1, 2, -2, -1},
			[][]int{},
		},
		"case8": testCase{
			[]int{-2, 0, 0, 2, 2},
			[][]int{
				{-2, 0, 2},
			},
		},
		"case9": testCase{
			[]int{-2, 0, 1, 1, 2},
			[][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		"case10": testCase{
			[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			[][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		"case11": testCase{
			[]int{-4, -2, -1},
			[][]int{},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := ThreeSum(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
