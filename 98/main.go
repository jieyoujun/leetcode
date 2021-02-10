package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 左根右
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validBST(root.Left, min, root.Val) && validBST(root.Right, root.Val, max)
}
