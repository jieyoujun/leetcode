package chongjianerchashulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == len(inorder) && len(preorder) > 0 {
		for i := range inorder {
			if inorder[i] == preorder[0] {
				return &TreeNode{
					Val:   preorder[0],
					Left:  BuildTree(preorder[1:i+1], inorder[0:i]),
					Right: BuildTree(preorder[i+1:], inorder[i+1:]),
				}
			}
		}
	}
	return nil
}
