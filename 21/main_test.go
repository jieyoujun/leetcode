package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		l1, l2, want *ListNode
	}
	testGroup := map[string]testCase{
		"case1": testCase{
			l1: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 4,
						Next: nil,
					},
				},
			},
			l2: &ListNode{Val: 1,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: nil,
					},
				},
			},
			want: &ListNode{Val: 1,
				Next: &ListNode{Val: 1,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 3,
							Next: &ListNode{Val: 4,
								Next: &ListNode{Val: 4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MergeTwoLists(tC.l1, tC.l2)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}