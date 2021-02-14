package linkedlistycleii

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	nfour := &ListNode{Val: -4}
	zero := &ListNode{Val: 0, Next: nfour}
	two := &ListNode{Val: 2, Next: zero}
	three := &ListNode{Val: 3, Next: two}
	nfour.Next = two
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: three,
			},
			want: two,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
