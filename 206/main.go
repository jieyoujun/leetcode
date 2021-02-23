package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// go语法糖
	// for head != nil {
	// 	head.Next, head, pre = pre, head.Next, head
	// }
	return pre
}
