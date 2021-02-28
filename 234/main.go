package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 找中点
	fast, slow := head, head
	var prev *ListNode
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 2. 翻转后半段
	prev.Next = nil
	var tail *ListNode
	for slow != nil {
		slow.Next, tail, slow = tail, slow, slow.Next
	}
	// 3. 比较前后半段
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true

	// 转数组
	// vals := []int{}
	// for p := head; p != nil; p = p.Next {
	// 	vals = append(vals, p.Val)
	// }
	// left, right := 0, len(vals)-1
	// for left < right {
	// 	if vals[left] != vals[right] {
	// 		return false
	// 	}
	// 	left++
	// 	right--
	// }
	// return true
}
