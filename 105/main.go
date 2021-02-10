package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root, rootIndexOfInorder := &TreeNode{Val: preorder[0]}, -1
	for idx, val := range inorder {
		if val == root.Val {
			rootIndexOfInorder = idx
			break
		}
	}
	root.Left = buildTree(preorder[1:rootIndexOfInorder+1], inorder[:rootIndexOfInorder])
	root.Right = buildTree(preorder[rootIndexOfInorder+1:], inorder[rootIndexOfInorder+1:])
	return root
}
