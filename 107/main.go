package binarytreelevelordertraversalii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result, nextLevel := [][]int{}, []*TreeNode{root}
	for len(nextLevel) > 0 {
		nnxtlvl := len(nextLevel)
		lvlVals := []int{}
		for i := 0; i < nnxtlvl; i++ {
			lvlVals = append(lvlVals, nextLevel[i].Val)
			if nextLevel[i].Left != nil {
				nextLevel = append(nextLevel, nextLevel[i].Left)
			}
			if nextLevel[i].Right != nil {
				nextLevel = append(nextLevel, nextLevel[i].Right)
			}
		}
		result = append([][]int{lvlVals}, result...)
		nextLevel = nextLevel[nnxtlvl:]
	}
	return result
}
