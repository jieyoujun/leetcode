package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 0 9 1 2 4/3 2 4
	// 3 2 4/0 9 1 2 4
	// O(m+n)
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB // a走完了接在b上
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA // b走完了接在a上
		}
	}
	return a

}
