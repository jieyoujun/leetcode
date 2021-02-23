package chongjianerchashulcof

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type testCase struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}
	testGroup := map[string]testCase{
		"case1": {[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := BuildTree(tC.preorder, tC.inorder)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
