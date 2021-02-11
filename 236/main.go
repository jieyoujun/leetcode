package lowestcommonancestorofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	oneInLeft := lowestCommonAncestor(root.Left, p, q)
	oneInRight := lowestCommonAncestor(root.Right, p, q)

	if oneInLeft == nil {
		return oneInRight
	}
	if oneInRight == nil {
		return oneInLeft
	}
	return root
}
