package erchashudeshendulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	if ld < rd {
		return rd + 1
	}
	return ld + 1
}
