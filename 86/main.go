package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ldummy, gdummy := ListNode{}, ListNode{}
	pld, pgd := &ldummy, &gdummy
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			pld.Next = &ListNode{Val: p.Val}
			pld = pld.Next
		} else {
			pgd.Next = &ListNode{Val: p.Val}
			pgd = pgd.Next
		}
	}
	pld.Next = gdummy.Next
	pgd.Next = nil
	return ldummy.Next
}
