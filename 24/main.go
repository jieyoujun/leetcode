package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = SwapPairs(next.Next)
	next.Next = head
	return next
}
