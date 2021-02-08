package reverselinkedlistii

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	dummy := ListNode{Next: head}
	mpre := &dummy

	for i := 0; i < m-1; i++ {
		mpre = mpre.Next
	}
	fmt.Println(mpre.Val)

	cur := mpre.Next
	for i := m; i < n; i++ {
		curNext := cur.Next
		cur.Next = curNext.Next
		curNext.Next = mpre.Next
		mpre.Next = curNext
	}
	return dummy.Next

	// 两趟扫描
	// var vals []int
	// for i, p := 0, head; p != nil; i, p = i+1, p.Next {
	// 	if i >= m-1 && i <= n-1 {
	// 		vals = append(vals, p.Val)
	// 	}
	// }
	// for i, p := 0, head; p != nil; i, p = i+1, p.Next {
	// 	if i >= m-1 && i <= n-1 {
	// 		p.Val = vals[len(vals)-1]
	// 		vals = vals[:len(vals)-1]
	// 	}
	// }

	return head
}
