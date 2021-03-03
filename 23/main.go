package mergeksortedlists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	nums := []int{}
	for _, l := range lists {
		for p := l; p != nil; p = p.Next {
			nums = append(nums, p.Val)
		}
	}
	sort.Ints(nums)
	dummy := &ListNode{}
	p := dummy
	for i := 0; i < len(nums); i++ {
		lnode := &ListNode{Val: nums[i]}
		p.Next = lnode
		p = p.Next
	}
	return dummy.Next
}
