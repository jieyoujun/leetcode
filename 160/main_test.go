package intersectionoftwolinkedlists

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	case1intersectionNode := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
						Next: case1intersectionNode,
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: case1intersectionNode,
						},
					},
				},
			},
			want: case1intersectionNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
