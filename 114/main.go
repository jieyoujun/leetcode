package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	left, right := root.Left, root.Right
	if left != nil {
		left = flatten(root.Left)
	}
	if right != nil {
		right = flatten(root.Right)
	}
	root.Left = nil
	if left != nil {
		l := left
		for {
			if l.Right == nil {
				l.Right = right
				break
			}
			l = l.Right
		}
		root.Right = left
	}
	return root
}
