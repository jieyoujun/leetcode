package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type testCase struct {
		head *ListNode
		n    int
		want *ListNode
	}
	testGroup := map[string]testCase{
		"case1": testCase{
			head: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5,
								Next: nil,
							},
						},
					},
				},
			},
			n: 2,
			want: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 5,
							Next: nil,
						},
					},
				},
			},
		},
		"case2": testCase{
			head: &ListNode{Val: 1,
				Next: nil,
			},
			n:    1,
			want: nil,
		},
		"case3": testCase{
			head: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: nil,
				},
			},
			n: 1,
			want: &ListNode{Val: 1,
				Next: nil,
			},
		},
		"case4": testCase{
			head: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: nil,
				},
			},
			n: 2,
			want: &ListNode{Val: 2,
				Next: nil,
			},
		},
	}

	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := RemoveNthFromEnd(tC.head, tC.n)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
