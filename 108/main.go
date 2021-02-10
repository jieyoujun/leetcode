package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[n/2]}
	root.Left = sortedArrayToBST(nums[:n/2])
	root.Right = sortedArrayToBST(nums[n/2+1:])
	return root
}
