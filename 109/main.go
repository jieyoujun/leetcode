package convertsortedlisttobinarysearchtree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// 定位中间节点 <- 快慢指针
	fast, slow := head, &ListNode{Next: head}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	root := &TreeNode{Val: slow.Next.Val}
	right := slow.Next.Next
	slow.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(right)
	return root
}
