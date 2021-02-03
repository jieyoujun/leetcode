package swapnodesinpairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	type testCase struct {
		head *ListNode
		want *ListNode
	}
	testGroup := map[string]testCase{
		"case1": testCase{
			head: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{Val: 2,
				Next: &ListNode{Val: 1,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 3,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := SwapPairs(tC.head)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
