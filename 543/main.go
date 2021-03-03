package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	ld := depth(root.Left, max)
	rd := depth(root.Right, max)
	if *max < ld+rd {
		*max = ld + rd
	}
	if ld < rd {
		return rd + 1
	}
	return ld + 1
}
