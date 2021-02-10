package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := 100000
	if root.Left != nil {
		ldepth := minDepth(root.Left)
		if depth > ldepth {
			depth = ldepth
		}
	}
	if root.Right != nil {
		rdepth := minDepth(root.Right)
		if depth > rdepth {
			depth = rdepth
		}
	}
	return depth + 1
}
