package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left-right >= -1 && left-right <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
