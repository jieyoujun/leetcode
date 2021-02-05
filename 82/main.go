package removeduplicatesfromsortedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{}
	pdummy := &dummy

	dup := false
	lastVal, head := head.Val, head.Next

	for head != nil {
		if head.Val == lastVal {
			dup = true
		} else {
			if !dup {
				pdummy.Next = &ListNode{Val: lastVal}
				pdummy = pdummy.Next
			}
			dup = false
			lastVal = head.Val
		}
		head = head.Next
	}
	if !dup {
		pdummy.Next = &ListNode{Val: lastVal}
		pdummy = pdummy.Next
	}
	return dummy.Next

	// 递归
	// if head.Val == head.Next.Val {
	// 	for head != nil && head.Next != nil && head.Val == head.Next.Val {
	// 		head = head.Next
	// 	}
	// 	return deleteDuplicates(head.Next)
	// }
	// head.Next = deleteDuplicates(head.Next)
	// return head
}
