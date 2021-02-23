package erchashuzhongheweimouyizhidelujinglcof

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	type testCase struct {
		root *TreeNode
		sum  int
		want [][]int
	}
	testGroup := map[string]testCase{
		"case1": {
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			22,
			[][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := PathSum(tC.root, tC.sum)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
