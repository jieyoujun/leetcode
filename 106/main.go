package constructbinarytreefrominorderandpostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	n := len(inorder)
	root, rootIndexOfInorder := &TreeNode{Val: postorder[n-1]}, -1
	for idx, val := range inorder {
		if val == root.Val {
			rootIndexOfInorder = idx
			break
		}
	}
	root.Left = buildTree(inorder[:rootIndexOfInorder], postorder[:rootIndexOfInorder])
	root.Right = buildTree(inorder[rootIndexOfInorder+1:], postorder[rootIndexOfInorder:n-1])
	return root
}
