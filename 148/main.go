package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// nlogn, 快排, 归并, 希尔, 堆
	// 快排
	// qsort(head, nil)
	// return head

	// 归并
	// 先一分为二
	fast, slow := head, &ListNode{Next: head}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	right := slow.Next
	slow.Next = nil
	left := head
	// 再归并排序
	right = sortList(right)
	left = sortList(left)
	dummy := &ListNode{}
	return mergeList(left, right, dummy)
}

func qsort(left, right *ListNode) {
	if left == right || left.Next == right {
		return
	}
	middle := gtMiddle(left, right)
	qsort(left, middle)
	qsort(middle.Next, right)
}

func gtMiddle(left, right *ListNode) *ListNode {
	if left == right || left.Next == right {
		return left
	}
	pivot := left.Val
	p1, p2 := left, left.Next
	for p2 != right {
		if p2.Val < pivot {
			p1 = p1.Next
			p1.Val, p2.Val = p2.Val, p1.Val
		}
		p2 = p2.Next
	}
	p1.Val, left.Val = left.Val, p1.Val
	// 小于pivot的放在p1前面
	return p1
}

func mergeList(left, right, dummy *ListNode) *ListNode {
	p := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	if left == nil {
		p.Next = right
	}
	if right == nil {
		p.Next = left
	}
	return dummy.Next
}
