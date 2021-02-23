package lianbiaozhongdaoshudikgejiedianlcof

import (
	"reflect"
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {
	type testCase struct {
		head *ListNode
		k    int
		want *ListNode
	}
	testGroup := map[string]testCase{
		"case1": {
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			2,
			&ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := GetKthFromEnd(tC.head, tC.k)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
