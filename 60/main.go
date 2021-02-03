package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	nNode, node := 1, head
	for node.Next != nil {
		node = node.Next
		nNode++
	}
	node.Next = head // 成环
	k %= nNode
	for i := 1; i <= nNode-k; i++ {
		node = node.Next
	}
	head, node.Next = node.Next, nil // 断开
	return head
}
