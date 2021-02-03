package lianbiaozhongdaoshudikgejiedianlcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	fast := head
	for i := 0; i < k-1; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
