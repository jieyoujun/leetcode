package uniquebinarysearchlTreeii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树：左<根<右
func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	return trees(1, n)
}

func trees(left, right int) []*TreeNode {
	lTree := []*TreeNode{}
	if left > right {
		lTree = append(lTree, nil)
		return lTree
	}
	for i := left; i <= right; i++ {
		leftTrees := trees(left, i-1)
		rightTrees := trees(i+1, right)

		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				lTree = append(lTree, &TreeNode{i, lt, rt})
			}
		}
	}
	return lTree
}
