package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fastNode, slowNode := head, head
	for i := 0; i < n; i++ {
		fastNode = fastNode.Next
	}
	if fastNode == nil {
		return head.Next
	}
	for fastNode.Next != nil {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	slowNode.Next = slowNode.Next.Next
	return head
}
