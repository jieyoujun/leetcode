package erweishuzuzhongdechazhaolcof

import (
	"reflect"
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
	type testCase struct {
		matrix [][]int
		target int
		want   bool
	}
	testGroup := map[string]testCase{
		"case1": {
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		"case2": {
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
		"case3": {
			[][]int{
				{-5},
			},
			-5,
			true,
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := FindNumberIn2DArray(tC.matrix, tC.target)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
