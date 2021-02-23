package uniquebinarysearchtreesii

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
	treeList := []*TreeNode{}
	if left > right {
		treeList = append(treeList, nil)
		return treeList
	}
	for i := left; i <= right; i++ {
		leftTrees := trees(left, i-1)
		rightTrees := trees(i+1, right)

		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				treeList = append(treeList, &TreeNode{i, lt, rt})
			}
		}
	}
	return treeList
}
