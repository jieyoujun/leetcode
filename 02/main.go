package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{Val: 0}
	curNode := ret
	n1, n2, carry := 0, 0, 0 // 进位
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		curNode.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		curNode = curNode.Next
		carry = (n1 + n2 + carry) / 10
	}
	return ret.Next
}

// 未考虑溢出
// func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil || l2 == nil {
// 		return &ListNode{}
// 	}
// 	sum := 0
// 	for i, node := 0, l1; node != nil; i, node = i+1, node.Next {
// 		sum += int(math.Pow10(i)) * node.Val
// 	}
// 	for i, node := 0, l2; node != nil; i, node = i+1, node.Next {
// 		sum += int(math.Pow10(i)) * node.Val
// 	}
// 	ret := &ListNode{}
// 	if sum == 0 {
// 		return ret
// 	}
// 	curNode := ret
// 	for sum > 0 {
// 		curNode.Next = &ListNode{Val: sum % 10}
// 		curNode = curNode.Next
// 		sum /= 10
// 	}
// 	return ret.Next
// }
